package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"
    "sync"
    "gopkg.in/yaml.v2"
    "github.com/julienschmidt/httprouter"
)

// Config structure to hold the configuration loaded from local.yml
type Config struct {
    Host                   string        `yaml:"host"`
    Port                   string        `yaml:"port"`
    CertFile               string        `yaml:"cert_file"`
    KeyFile                string        `yaml:"key_file"`
    StaticDir              string        `yaml:"static_dir"`
    StaticPath             string        `yaml:"static_path"`
    HealthPath             string        `yaml:"health_path"`
    RequestTimeout         time.Duration `yaml:"request_timeout"`
    Default404Handler      bool          `yaml:"default_404_handler"`
    HandleMethodNotAllowed bool          `yaml:"handle_method_not_allowed"`
    DefaultContentType     string        `yaml:"default_content_type"`
}

// Todo struct to represent a to-do item
type Todo struct {
    ID   string `json:"id"`
    Task string `json:"task"`
    Done bool   `json:"done"`
}

// HTTPServer struct
type HTTPServer struct {
    config *Config
    router *httprouter.Router
    server *http.Server
    todos  map[string]Todo
    mu     sync.Mutex
}

// Load configuration from local.yml with default values
func loadConfigWithDefaults(configFile string) (*Config, error) {
    data, err := ioutil.ReadFile(configFile)
    if err != nil {
        return nil, err
    }

    config := &Config{
        Host:                   "0.0.0.0",
        Port:                   "8080",
        CertFile:               "",
        KeyFile:                "",
        StaticDir:              "./static",
        StaticPath:             "/static/",
        HealthPath:             "/health",
        RequestTimeout:         5 * time.Second,
        Default404Handler:      true,
        HandleMethodNotAllowed: false,
        DefaultContentType:     "application/json",
    }

    err = yaml.Unmarshal(data, config)
    if err != nil {
        return nil, err
    }
    return config, nil
}

// NewHTTPServer creates a new HTTP server instance
func NewHTTPServer(config *Config) *HTTPServer {
    server := &HTTPServer{
        config: config,
        router: httprouter.New(),
        todos:  make(map[string]Todo),
    }

    if !strings.HasSuffix(config.StaticPath, "/") && len(config.StaticPath) > 0 {
        config.StaticPath += "/"
    }

    server.server = &http.Server{
        Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
        Handler: server.router,
    }

    return server
}

// Middleware for authentication
func (s *HTTPServer) authenticationMiddleware(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        username, password, ok := r.BasicAuth()
        if !ok || username != "admin" || password != "password" {
            s.writeMessage(401, "Unauthorized", r, w, nil)
            return
        }
        next(w, r, p)
    }
}

// ServeFiles serves static files
func (s *HTTPServer) ServeFiles(path, prefix string, root http.FileSystem) {
    if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
        panic("path must end with /*filepath")
    }
    fileServer := http.FileServer(root)
    s.router.GET(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        if len(prefix) != 0 {
            if strings.HasSuffix(prefix, "/") {
                r.URL.Path = prefix + p.ByName("filepath")
            } else {
                r.URL.Path = prefix + "/" + p.ByName("filepath")
            }
        } else {
            r.URL.Path = p.ByName("filepath")
        }
        fileServer.ServeHTTP(w, r)
    })
}

// Setup server routes and handlers
func (s *HTTPServer) Setup() {
    if s.config.Default404Handler {
        s.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            s.writeMessage(404, "Page not found", r, w, nil)
        })
    }
    s.router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
        s.writeMessage(500, "Internal server error", r, w, nil)
    }

    if len(s.config.StaticDir) != 0 {
        fi, e := os.Stat(s.config.StaticDir)
        if e != nil {
            log.Println("Error:", e.Error())
        } else {
            if fi.IsDir() {
                s.ServeFiles(s.config.StaticPath+"*filepath", s.config.StaticPath, http.Dir(s.config.StaticDir))
            } else {
                log.Println("Provided static_dir is not a directory")
            }
        }
    }

    if len(s.config.HealthPath) != 0 {
        s.router.GET(s.config.HealthPath, func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
            w.WriteHeader(http.StatusOK)
            log.Println("HEALTH OK")
        })
    }

    // Add to-do handlers
    s.router.GET("/todos", s.authenticationMiddleware(s.getTodos))
    s.router.POST("/todos", s.authenticationMiddleware(s.createTodo))
    s.router.GET("/todos/:id", s.authenticationMiddleware(s.getTodo))
    s.router.PUT("/todos/:id", s.authenticationMiddleware(s.updateTodo))
    s.router.DELETE("/todos/:id", s.authenticationMiddleware(s.deleteTodo))
}

// WriteMessage sends a formatted JSON response
func (s *HTTPServer) writeMessage(statusCode int, defaultMessage string, request *http.Request, writer http.ResponseWriter, errLocal error) {
    if errLocal != nil {
        log.Println("Error:", errLocal.Error())
    }
    writer.Header().Add("Content-Type", s.config.DefaultContentType)
    writer.WriteHeader(statusCode)
    message := fmt.Sprintf(`{ "code": "SE_%d", "lang": "en", "message": "%d ERROR", "data": {} }`, statusCode, statusCode)
    if _, err := writer.Write([]byte(message)); err != nil {
        log.Println("Write Error:", err.Error())
    }
}

// Handler functions for /todos endpoints
func (s *HTTPServer) getTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    s.mu.Lock()
    defer s.mu.Unlock()
    todos := []Todo{}
    for _, todo := range s.todos {
        todos = append(todos, todo)
    }
    s.writeJSON(w, http.StatusOK, todos)
}

func (s *HTTPServer) createTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        s.writeMessage(400, "Bad Request", r, w, err)
        return
    }

    s.mu.Lock()
    defer s.mu.Unlock()
    s.todos[todo.ID] = todo
    s.writeJSON(w, http.StatusCreated, todo)
}

func (s *HTTPServer) getTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    s.mu.Lock()
    defer s.mu.Unlock()
    todo, exists := s.todos[id]
    if !exists {
        s.writeMessage(404, "Not Found", r, w, nil)
        return
    }
    s.writeJSON(w, http.StatusOK, todo)
}

func (s *HTTPServer) updateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    var updatedTodo Todo
    if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
        s.writeMessage(400, "Bad Request", r, w, err)
        return
    }

    s.mu.Lock()
    defer s.mu.Unlock()
    todo, exists := s.todos[id]
    if !exists {
        s.writeMessage(404, "Not Found", r, w, nil)
        return
    }
    todo.Task = updatedTodo.Task
    todo.Done = updatedTodo.Done
    s.todos[id] = todo
    s.writeJSON(w, http.StatusOK, todo)
}

func (s *HTTPServer) deleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    s.mu.Lock()
    defer s.mu.Unlock()
    _, exists := s.todos[id]
    if !exists {
        s.writeMessage(404, "Not Found", r, w, nil)
        return
    }
    delete(s.todos, id)
    w.WriteHeader(http.StatusNoContent)
}

// WriteJSON sends a JSON response
func (s *HTTPServer) writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Println("JSON Encoding Error:", err.Error())
    }
}

// Start starts the HTTP server with or without TLS
func (s *HTTPServer) Start() error {
    
    log.Printf("HTTP server started at %s:%s", s.config.Host, s.config.Port)
    if len(s.config.CertFile) != 0 && len(s.config.KeyFile) != 0 {
        return s.server.ListenAndServeTLS(s.config.CertFile, s.config.KeyFile)
    } else {
        return s.server.ListenAndServe()
    }
}

// Graceful shutdown handling
func (s *HTTPServer) GracefulShutdown(wait time.Duration, wg *sync.WaitGroup) {
    defer wg.Done()
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

    <-quit
    log.Println("Shutting down server...")
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    if err := s.server.Shutdown(ctx); err != nil {
        log.Fatalf("Server Shutdown Failed: %v", err)
    }
    log.Println("Server gracefully stopped")
}

func main() {
    // Load configuration
    config, err := loadConfigWithDefaults("local.yml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Create and setup server
    server := NewHTTPServer(config)
    server.Setup()

    // Wait group for graceful shutdown
    var wg sync.WaitGroup
    wg.Add(1)
    go server.GracefulShutdown(10*time.Second, &wg)

    // Start the server in a separate goroutine
    go func() {
        if err := server.Start(); err != nil {
            log.Fatalf("Server failed to start: %v", err)
        }
    }()

    // Wait for graceful shutdown
    wg.Wait()
}

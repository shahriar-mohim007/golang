package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// HTTPServer holds configuration for the HTTP server
type HTTPServer struct {
	host                 string
	port                 string
	certFile              string
	keyFile               string
	staticDir             string
	staticPath            string
	healthPath            string
	requestTimeout        time.Duration
	default404Handler     bool
	handleMethodNotAllowed bool
	defaultContentType    string
	router                *httprouter.Router
}

// NewHTTPServer creates a new instance of HTTPServer
func NewHTTPServer(host, port, certFile, keyFile, staticDir, staticPath, healthPath string, requestTimeout time.Duration, default404Handler, handleMethodNotAllowed bool, defaultContentType string) *HTTPServer {
	server := &HTTPServer{
		host:                 host,
		port:                 port,
		certFile:             certFile,
		keyFile:              keyFile,
		staticDir:            staticDir,
		staticPath:           staticPath,
		healthPath:           healthPath,
		requestTimeout:       requestTimeout,
		default404Handler:    default404Handler,
		handleMethodNotAllowed: handleMethodNotAllowed,
		defaultContentType:   defaultContentType,
		router:               httprouter.New(),
	}

	if !strings.HasSuffix(server.staticPath, "/") && len(server.staticPath) > 0 {
		server.staticPath = server.staticPath + "/"
	}

	return server
}

func (s *HTTPServer) buildDefaultMessage(code uint32) string {
	return fmt.Sprintf(`
		{
			"code": "SE_%d",
			"lang": "en",
			"message": "%d ERROR",
			"data": {}
		}
	`, code, code)
}

func (s *HTTPServer) getMessage(key, defaultValue, lang string) string {
	// Placeholder for fetching localized messages
	return fmt.Sprintf(`{"message": "%s"}`, defaultValue)
}

func (s *HTTPServer) getLanguage(r *http.Request) string {
	l := r.Header.Get("Accept-Language")
	if len(l) == 0 {
		l = "en"
	}
	return l
}

func (s *HTTPServer) writeMessage(statusCode int, defaultMessage string, request *http.Request, writer http.ResponseWriter, errLocal error) {
	if errLocal != nil {
		logger.L().Error(errLocal.Error())
	}

	writer.Header().Add("Content-Type", s.defaultContentType)
	writer.WriteHeader(statusCode)
	if _, err := writer.Write([]byte(s.getMessage(fmt.Sprintf("s%dm", statusCode), defaultMessage, s.getLanguage(request)))); err != nil {
		logger.L().Error(err.Error())
	}
}

func (s *HTTPServer) s401m(request *http.Request, writer http.ResponseWriter, errLocal error) {
	s.writeMessage(401, s.buildDefaultMessage(401), request, writer, errLocal)
}

func (s *HTTPServer) s403m(request *http.Request, writer http.ResponseWriter, errLocal error) {
	s.writeMessage(403, s.buildDefaultMessage(403), request, writer, errLocal)
}

func (s *HTTPServer) s404m(request *http.Request, writer http.ResponseWriter, errLocal error) {
	s.writeMessage(404, s.buildDefaultMessage(404), request, writer, errLocal)
}

func (s *HTTPServer) s405m(request *http.Request, writer http.ResponseWriter, errLocal error) {
	s.writeMessage(405, s.buildDefaultMessage(405), request, writer, errLocal)
}

func (s *HTTPServer) s500m(request *http.Request, writer http.ResponseWriter, errLocal error) {
	s.writeMessage(500, s.buildDefaultMessage(500), request, writer, errLocal)
}

func (s *HTTPServer) debugMessage(request *http.Request) {
	logger.L().Debug("request local timeout in seconds", zap.Duration("timeout", s.requestTimeout))
	logger.L().Debug("request started")
	logger.L().Debug("request data",
		zap.String("path", request.URL.Path),
		zap.String("method", request.Method),
		zap.String("path_with_query", request.RequestURI))
}

// Middleware for authentication
func (s *HTTPServer) authenticationMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Example basic authentication
		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != "password" {
			s.s401m(r, w, nil)
			return
		}

		next(w, r, p)
	}
}

func (s *HTTPServer) ServeFiles(path string, prefix string, root http.FileSystem) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path must end with /*filepath in path '" + path + "'")
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

func (s *HTTPServer) Setup() {
	if s.default404Handler {
		s.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.debugMessage(r)
			s.s404m(r, w, nil)
		})
	}

	s.router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		s.debugMessage(r)
		s.s500m(r, w, nil)
	}

	// Serve static files
	if len(s.staticDir) != 0 {
		fi, e := os.Stat(s.staticDir)
		if e != nil {
			logger.L().Error(e.Error())
		} else {
			if fi.IsDir() {
				s.ServeFiles(s.staticPath+"*filepath", s.staticPath, http.Dir(s.staticDir))
			} else {
				logger.L().Error("provided static_dir in the manifest conf is not directory")
			}
		}
	}

	// Health check route
	if len(s.healthPath) != 0 {
		s.router.GET(s.healthPath, func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
			w.WriteHeader(http.StatusOK)
			logger.L().Info("HEALTH OK")
		})
	}

	logger.L().Info("HTTP server setup complete", zap.String("host", s.host), zap.String("port", s.port))
}

func (s *HTTPServer) Start() error {
	logger.L().Info("HTTP server started at " + s.host + ":" + s.port)

	if len(s.certFile) != 0 && len(s.keyFile) != 0 {
		return http.ListenAndServeTLS(s.host+":"+s.port, s.certFile, s.keyFile, s.router)
	} else {
		return http.ListenAndServe(s.host+":"+s.port, s.router)
	}
}

func main() {
	// Example configuration
	config := model.ConfigMap{
		"host":                    "0.0.0.0",
		"port":                    "8080",
		"cert_file":               "",
		"key_file":                "",
		"static_dir":              "./static",
		"static_path":             "/static/",
		"health_path":             "/health",
		"default_request_timeout": "1s",
		"default_404_handler_enabled": "true",
		"handle_method_not_allowed": "false",
		"default_content_type":    "application/json",
	}

	server := NewHTTPServer(
		config.String("host", "0.0.0.0"),
		config.String("port", "8080"),
		config.String("cert_file", ""),
		config.String("key_file", ""),
		config.String("static_dir", ""),
		config.String("static_path", "/static/"),
		config.String("health_path", ""),
		config.Duration("default_request_timeout", time.Second),
		config.Bool("default_404_handler_enabled", true),
		config.Bool("handle_method_not_allowed", false),
		config.String("default_content_type", "application/json"),
	)

	server.Setup()

	// Adding handlers
	server.router.GET("/example", server.authenticationMiddleware(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Example route"))
	}))

	if err := server.Start(); err != nil {
		logger.L().Fatal("Server failed to start", zap.Error(err))
	}
}

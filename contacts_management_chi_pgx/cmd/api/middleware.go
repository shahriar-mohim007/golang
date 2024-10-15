package main

import (
	utils "chi_pgx/utils"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type contextKey string

const userContextKey = contextKey("userID")

func (app *application) AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := extractTokenFromHeader(r)
		if tokenStr == "" {
			app.logger.PrintError(fmt.Errorf("no token provided"), map[string]string{
				"context": "authorization",
			})
			_ = Unauthorized.WriteToResponse(w, nil)
			return
		}

		var claims utils.Claims
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(app.config.SecretKey), nil
		})

		if err != nil || !token.Valid {
			app.logger.PrintError(fmt.Errorf("invalid token"), map[string]string{
				"context": "authorization",
			})
			_ = Unauthorized.WriteToResponse(w, nil)
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, claims.UserID.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func (app *application) RateLimitMiddleware(next http.Handler) http.Handler {

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.config.LimiterEnabled {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				app.logger.PrintError(err, map[string]string{
					"context": "host port split error",
				})
				_ = Unauthorized.WriteToResponse(w, nil)
				return
			}

			mu.Lock()
			if _, found := clients[ip]; !found {
				clients[ip] = &client{
					limiter: rate.NewLimiter(
						rate.Limit(app.config.Rps),
						app.config.Burst),
				}
			}

			clients[ip].lastSeen = time.Now()

			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				_ = RateLimitExceeded.WriteToResponse(w, nil)
				return
			}

			mu.Unlock()
		}

		next.ServeHTTP(w, r)
	})

}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(authHeader, "Bearer ")
}

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userContextKey).(string)
	return userID, ok
}

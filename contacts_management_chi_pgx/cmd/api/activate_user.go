package main

import (
	utils "chi_pgx/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func (app *application) handleActivateUser(w http.ResponseWriter, req *http.Request) {

	tokenString := req.URL.Query().Get("token")
	ctx := req.Context()

	if tokenString == "" {
		app.logger.PrintError(fmt.Errorf("missing token"), map[string]string{
			"context": "missing token",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	var claims utils.Claims
	secretKey := app.config.SecretKey

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		app.logger.PrintError(err, map[string]string{
			"context": "invalid token",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	userID := claims.UserID

	err = app.Repository.ActivateUserByID(ctx, userID)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "failed to activate user",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}
	_ = UserActivated.WriteToResponse(w, nil)
	return
}

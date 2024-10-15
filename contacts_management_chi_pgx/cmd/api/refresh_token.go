package main

import (
	utils "chi_pgx/utils"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type RefreshRequestPayload struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponsePayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (app *application) handleRefreshToken(w http.ResponseWriter, req *http.Request) {

	refreshRequest := RefreshRequestPayload{}
	err := json.NewDecoder(req.Body).Decode(&refreshRequest)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Invalid JSON",
		})
		_ = ValidDataNotFound.WriteToResponse(w, nil)
		return
	}

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(refreshRequest.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.config.SecretKey), nil
	})

	if err != nil || !token.Valid {
		app.logger.PrintError(err, map[string]string{
			"context": "Invalid token",
		})
		_ = Unauthorized.WriteToResponse(w, nil)
		return
	}

	userID, err := uuid.FromString(claims.Subject)

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Error parsing UUID",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	accessToken, err := utils.GenerateJWT(userID, utils.ScopeAuthentication, app.config.SecretKey, 2*time.Hour)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Error generating access token",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	newRefreshToken, err := utils.GenerateRefreshToken(claims.Subject, app.config.SecretKey)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Error generating refresh token",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	tokenResponse := RefreshResponsePayload{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}

	_ = loginSuccess.WriteToResponse(w, tokenResponse)
	return

}

package main

import (
	"chi_pgx/internal/domain"
	utils "chi_pgx/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type requestPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responsePayload struct {
	ActivateToken string `json:"activate_token"`
}

func (app *application) handleRegisterUser(w http.ResponseWriter, req *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			app.logger.PrintError(fmt.Errorf("recovered from panic: %v", r), map[string]string{
				"context": "panic recovery",
			})
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	request := requestPayload{}
	ctx := req.Context()
	err := json.NewDecoder(req.Body).Decode(&request)

	if err != nil || request.Name == "" || request.Email == "" || request.Password == "" {

		app.logger.PrintError(err, map[string]string{
			"context": "Invalid input",
		})
		_ = ValidDataNotFound.WriteToResponse(w, nil)
		return
	}
	user, err := app.Repository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.logger.PrintError(fmt.Errorf("user with email %s not found", request.Email), map[string]string{
				"context": "user lookup",
				"email":   request.Email,
			})
		} else {
			app.logger.PrintError(err, map[string]string{
				"context": "Error fetching user by email",
			})
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	if user != nil {
		app.logger.PrintInfo(fmt.Sprintf("User already exists: %s", request.Email), map[string]string{
			"context": "user registration",
		})
		_ = UserAlreadyExist.WriteToResponse(w, nil)
		return
	}

	passwordHash, err := utils.HashPassword(request.Password)

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Failed to hash password",
		})

		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userID, err := uuid.NewV4()

	if err != nil {
		http.Error(w, "Error Generating Id", http.StatusInternalServerError)
		return
	}

	user = &domain.User{
		ID:       userID,
		Name:     request.Name,
		Email:    request.Email,
		Password: passwordHash,
		IsActive: false,
	}

	if err = app.Repository.CreateUser(ctx, user); err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Failed to create user",
		})
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	ttl := 2 * time.Hour

	token, err := utils.GenerateJWT(user.ID, utils.ScopeActivation, app.config.SecretKey, ttl)

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Failed to activation token",
		})

		http.Error(w, "Error creating activation token", http.StatusInternalServerError)
		return
	}

	response := responsePayload{
		ActivateToken: token,
	}

	_ = UserCreated.WriteToResponse(w, response)
	return

}

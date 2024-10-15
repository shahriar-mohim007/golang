package main

import (
	"chi_pgx/internal/domain"
	utils "chi_pgx/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type RegistrationRequestPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegistrationResponsePayload struct {
	ActivateToken string `json:"activate_token"`
}

func (app *application) handleRegisterUser(w http.ResponseWriter, req *http.Request) {

	request := RegistrationRequestPayload{}
	ctx := req.Context()
	err := json.NewDecoder(req.Body).Decode(&request)

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Invalid JSON",
		})
		_ = ValidDataNotFound.WriteToResponse(w, nil)
		return
	}
	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Invalid payload",
		})
		_ = ValidDataNotFound.WriteToResponse(w, nil)
		return
	}

	user, err := app.Repository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.logger.PrintInfo(fmt.Sprintf("user with email %s not found", request.Email), map[string]string{})
		} else {
			app.logger.PrintError(err, map[string]string{
				"context": "Error fetching user by email",
			})
			_ = InternalError.WriteToResponse(w, nil)
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

		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	userID, err := uuid.NewV4()

	if err != nil {
		_ = InternalError.WriteToResponse(w, nil)
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
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	ttl := 2 * time.Hour

	token, err := utils.GenerateJWT(user.ID, utils.ScopeActivation, app.config.SecretKey, ttl)

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Failed to activation token",
		})

		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	response := RegistrationResponsePayload{
		ActivateToken: token,
	}

	_ = UserCreated.WriteToResponse(w, response)
	return

}

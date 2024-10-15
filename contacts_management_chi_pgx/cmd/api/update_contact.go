package main

import (
	"chi_pgx/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"net/http"
)

func (app *application) handleUpdateContact(w http.ResponseWriter, req *http.Request) {
	contactID := chi.URLParam(req, "id")
	uuidContactID, err := uuid.FromString(contactID)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Error parsing contact ID",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}
	ctx := req.Context()

	contact, err := app.Repository.GetContactByID(ctx, uuidContactID)
	if err != nil {
		_ = NotFound.WriteToResponse(w, nil)
		return
	}

	requestPayload := ContactRequestPayload{}
	err = json.NewDecoder(req.Body).Decode(&requestPayload)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Invalid JSON",
		})
		_ = ValidDataNotFound.WriteToResponse(w, nil)
		return
	}

	if requestPayload.Phone != "" {
		contact.Phone = requestPayload.Phone
	}
	if requestPayload.Street != "" {
		contact.Street = requestPayload.Street
	}
	if requestPayload.City != "" {
		contact.City = requestPayload.City
	}
	if requestPayload.State != "" {
		contact.State = requestPayload.State
	}
	if requestPayload.ZipCode != "" {
		contact.ZipCode = requestPayload.ZipCode
	}
	if requestPayload.Country != "" {
		contact.Country = requestPayload.Country
	}

	updatedContact := domain.Contact{
		Phone:   contact.Phone,
		Street:  contact.Street,
		City:    contact.City,
		State:   contact.State,
		ZipCode: contact.ZipCode,
		Country: contact.Country,
	}

	err = app.Repository.PatchContact(ctx, uuidContactID, &updatedContact)
	if err != nil {
		_ = InternalError.WriteToResponse(w, err)
		return
	}
	response := ContactResponse{
		ID:      contactID,
		Phone:   contact.Phone,
		Street:  contact.Street,
		City:    contact.City,
		State:   contact.State,
		ZipCode: contact.ZipCode,
		Country: contact.Country,
	}

	_ = ContactUpdated.WriteToResponse(w, response)
	return
}

package main

import (
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"net/http"
	"strings"
)

func (app *application) handleGetContact(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	contactID, err := uuid.FromString(id)

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Error creating contact",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}
	ctx := req.Context()
	contact, err := app.Repository.GetContactByID(ctx, contactID)
	if err != nil {
		if strings.Contains(err.Error(), "no contact found") {
			_ = NotFound.WriteToResponse(w, nil)
		} else {
			app.logger.PrintError(err, map[string]string{
				"context": "Error fetching contact",
			})
			_ = InternalError.WriteToResponse(w, nil)
		}
		return
	}

	_ = ContactRetrieved.WriteToResponse(w, contact)
	return
}

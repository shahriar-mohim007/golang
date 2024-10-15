package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"net/http"
)

func (app *application) handleDeleteContact(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	contactID, err := uuid.FromString(id)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"context": "Error parsing contact ID",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	ctx := req.Context()
	err = app.Repository.DeleteContactByID(ctx, contactID)

	if err != nil {
		if err == sql.ErrNoRows {
			app.logger.PrintError(err, map[string]string{
				"context": "Contact not found",
			})
			_ = InternalError.WriteToResponse(w, nil)

		} else {
			app.logger.PrintError(err, map[string]string{
				"context": "Error deleting contact",
			})
			_ = InternalError.WriteToResponse(w, nil)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

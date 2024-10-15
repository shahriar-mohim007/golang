package main

import (
	"github.com/gofrs/uuid"
	"net/http"
)

type ContactResponse struct {
	ID      string `json:"id"`
	Phone   string `json:"phone"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type ContactsResponse struct {
	Contacts []ContactResponse `json:"contacts"`
}

func (app *application) handleGetAllContact(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	userID, _ := GetUserIDFromContext(ctx)
	uuID, err := uuid.FromString(userID)

	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"Context": "Error parsing UUID",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	contacts, err := app.Repository.GetAllContacts(ctx, uuID)
	if err != nil {
		app.logger.PrintError(err, map[string]string{
			"Context": "Error fetching contacts",
		})
		_ = InternalError.WriteToResponse(w, nil)
		return
	}

	var contactResponses []ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, ContactResponse{
			ID:      contact.ID.String(),
			Phone:   contact.Phone,
			Street:  contact.Street,
			City:    contact.City,
			State:   contact.State,
			ZipCode: contact.ZipCode,
			Country: contact.Country,
		})
	}

	response := ContactsResponse{
		Contacts: contactResponses,
	}

	_ = ContactRetrieved.WriteToResponse(w, response)
	return
}

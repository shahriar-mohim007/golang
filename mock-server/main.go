package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	Token           string `json:"token"`
	TokenExpireTime string `json:"tokenexpiretime"`
	DataSign        string `json:"datasign"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type SendMoneyRequestData struct {
	RegisterRef   string `json:"registerref"`
	ReceiverAcc   string `json:"receiverac"`
	ReceiverTitle string `json:"receivertitle"`
	Amount        int    `json:"amount"`
	DataSign      string `json:"datasign"`
}

type SendMoneyRequestResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    SendMoneyRequestData `json:"data"`
}

type ConfirmTransactionData struct {
	TransactionStatus    string `json:"transactionstatus"`
	BankTransactionID    string `json:"banktransactionid"`
	ClientReferenceID    string `json:"clientreferenceid"`
	Amount               int    `json:"amount"`
	TransactionTimestamp string `json:"transactiontimestamp"`
	DataSign             string `json:"datasign"`
}

type ConfirmTransactionResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    ConfirmTransactionData `json:"data"`
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")
	if apiKey != "24E876DA7839C332FD647B393B48EBDA8B108DCB" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	response := Response{
		Status:  "200",
		Message: "Success",
		Data: Data{
			Token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ4LXBhcnRpY3VsYXItMSI6ImhOb256MDk1QTlCWTlXbkxGRmhTZnZlZTJqWUVqdW1KanhXZTloS0lIMzVTODdQbFR6L3ZPVHA3WnhWMUc2TEk0bnBSaEFRekhrWGVDdmNTSG9maXdwS3gvVnNQRzFmYTJRQnRLVllVSWtc",
			TokenExpireTime: "13/02/2024 13:41:20",
			DataSign:        "VHA3WnhWMUc2TEk0bnBSaEFRekhrWGVDdmNTSG9maXdwS3gvVnNQRzFmYTJRQnRLVllVSWtc",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func sendMoneyRequestHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")
	if apiKey != "24E876DA7839C332FD647B393B48EBDA8B108DCB" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	token, err := extractToken(r)
	if err != nil || token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ4LXBhcnRpY3VsYXItMSI6ImhOb256MDk1QTlCWTlXbkxGRmhTZnZlZTJqWUVqdW1KanhXZTloS0lIMzVTODdQbFR6L3ZPVHA3WnhWMUc2TEk0bnBSaEFRekhrWGVDdmNTSG9maXdwS3gvVnNQRzFmYTJRQnRLVllVSWtc" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	response := SendMoneyRequestResponse{
		Status:  "200",
		Message: "Success",
		Data: SendMoneyRequestData{
			RegisterRef:   "1fc83e28-09af-4f09-b277-d209047d0a9f",
			ReceiverAcc:   "3555101000000000",
			ReceiverTitle: "mr. xyz hasan",
			Amount:        500,
			DataSign:      "kvM8a8z0D4nRjOt0u3LklIk70Bq1JvrYD84gZw6Xy9mh1WsQLJmag9OqXHkdmJDpw1hH8tYOczOCtm",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func confirmTransactionHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")
	if apiKey != "24E876DA7839C332FD647B393B48EBDA8B108DCB" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	token, err := extractToken(r)
	if err != nil || token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ4LXBhcnRpY3VsYXItMSI6ImhOb256MDk1QTlCWTlXbkxGRmhTZnZlZTJqWUVqdW1KanhXZTloS0lIMzVTODdQbFR6L3ZPVHA3WnhWMUc2TEk0bnBSaEFRekhrWGVDdmNTSG9maXdwS3gvVnNQRzFmYTJRQnRLVllVSWtc" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	response := ConfirmTransactionResponse{
		Status:  "200",
		Message: "Success",
		Data: ConfirmTransactionData{
			TransactionStatus:    "success",
			BankTransactionID:    "20220103-7-20001",
			ClientReferenceID:    "RXYPCRNT",
			Amount:               500,
			TransactionTimestamp: "13/02/2024 13:52:20",
			DataSign:             "AoNUr/z20Lxi9JEbc2Qj+SXme9hdyA3DawVVHt+SXme9hd6Eon9w+Lj2h+Lj2h+ozYyXdGf++ozYyXdGf+",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/v1/auth", authHandler)
	http.HandleFunc("/api/v1/sendmoneyrequest", sendMoneyRequestHandler)
	http.HandleFunc("/api/v1/confirmtransaction", confirmTransactionHandler)

	fmt.Println("Starting mock server on :8080")
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

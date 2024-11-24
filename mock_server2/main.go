package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type AuthResponse struct {
	ErrorCode  string `json:"errorcode"`
	ErrorMsg   string `json:"errormsg"`
	Time       string `json:"timestamp"`
	Token      string `json:"token"`
	TokeExpire string `json:"tokenExpireDateTime"`
}

type SendAccountResponse struct {
	ErrorCode string `json:"errorcode"`
	ErrorMsg  string `json:"errormsg"`
	Time      string `json:"timestamp"`
	SeblId    string `json:"seblTransactionId"`
}

func main() {
	http.HandleFunc("/requesttoken", func(w http.ResponseWriter, r *http.Request) {
		// Mock response for authentication
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var authPayload map[string]string
		if err := json.NewDecoder(r.Body).Decode(&authPayload); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		// Mock a successful response
		authResponse := AuthResponse{
			ErrorCode:  "0",
			ErrorMsg:   "SUCCESS",
			Time:       "2024-11-20T12:00:00Z",
			Token:      "mocked_token_123",
			TokeExpire: "2024-11-21T12:00:00Z",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authResponse)
	})

	http.HandleFunc("/sendtoaccount", func(w http.ResponseWriter, r *http.Request) {
		// Mock response for sending money to an account
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var sendPayload map[string]string
		if err := json.NewDecoder(r.Body).Decode(&sendPayload); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		// Mock a successful transaction
		sendResponse := SendAccountResponse{
			ErrorCode: "0",
			ErrorMsg:  "SUCCESS",
			Time:      "2024-11-20T12:05:00Z",
			SeblId:    "mocked_sebl_id_456",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sendResponse)
	})

	log.Println("Mock Southeast Bank server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

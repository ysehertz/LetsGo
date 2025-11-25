package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"native-web-demo/internal/models"
)

// Home 是一个处理器函数。
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Native Go")
}

// UsersHandler 处理 /users 路径的请求。
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users := []models.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		}

		jsonResponse, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Error generating JSON", http.StatusInternalServerError)
			return // Add return here
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	case http.MethodPost:
		var newUser models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		log.Printf("Received new user: %+v", newUser) // %+v 会打印结构体字段名和值

		jsonResponse, err := json.Marshal(newUser) // Fix typo: jsonResponce -> jsonResponse
		if err != nil {
			http.Error(w, "Error generating JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

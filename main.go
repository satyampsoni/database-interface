package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// User represents a user in the database.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func getUsers(db *sql.DB) ([]User, error) {
	// ... Querying the database to retrieve users ...

	users := []User{
		{ID: 1, Username: "satyampsoni", Email: "satyampsoni@gmail.com"},
		{ID: 2, Username: "shivampsoni", Email: "shivampsoni@gmail.com"},
		{ID: 3, Username: "Armada", Email: "armada@gmail.com"},
	}

	return users, nil
}

func main() {
	// Establishing the database connection
	connStr := "host=localhost port=5432 user=satyampsoni password=Satyam!12 dbname=databaseinterface sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// HTTP handler function for retrieving users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		users, err := getUsers(db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// JSON encode the users and send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

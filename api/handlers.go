package api

import (
	"encoding/json"
	"log"
	"net/http"

	mdb "github.com/Whadislov/TTCompanion2/internal/my_db"
	mf "github.com/Whadislov/TTCompanion2/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion2/internal/my_types"

	"github.com/google/uuid"
)

// Handler for loading the database
func loadUserDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API : Received request to load user DB")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Header.Get("User-ID")
	if userID == "" {
		http.Error(w, "User is unidentified", http.StatusUnauthorized)
		return
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "Failed to parse user ID", http.StatusUnauthorized)
		return
	}

	mdb.SetUserIDOfSession(id)
	db, err := mdb.LoadDB()
	if err != nil {
		http.Error(w, "Failed to connect to database.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

// Handler for saving the local changes to the database
func saveDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API : Received request to save user DB")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var db mt.Database
	err := json.NewDecoder(r.Body).Decode(&db)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	err = mdb.SaveDB(&db)
	if err != nil {
		http.Error(w, "Failed to save database", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler to check if the API is ready to take requests
func IsApiReadyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is ready !"))
}

// loginHandler process the request to analyse the credentials, returns a Credential token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		sendJSONError(w, "Invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}

	userID, err := checkUserCredentials(creds.Username, creds.Password)
	if err != nil {
		sendJSONError(w, "Invalid username or password", "INVALID_USERNAME_OR_PASSWORD", http.StatusUnauthorized)
		return
	}

	credToken, err := generateJWT(userID)
	if err != nil {
		sendJSONError(w, "Could not generate credToken", "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	// Create a new session with gorilla. Session name = persistency-session, the name needs to be static
	log.Println("API : creation of a session")
	createSession(w, r, credToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"cred_token": credToken})
}

// logoutHandler deletes the session and the cookie of a user when he logs out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API : Received request to log user out")
	err := deleteSession(w, r)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

// signUpHandler process the request to create a new user, returns a Credential token
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API : Received request to sign user up")
	var signUpData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&signUpData); err != nil {
		sendJSONError(w, "Invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}

	// Load the whole database to register the new user. Could be optimised to request directly postgres
	db, err := mdb.LoadUsersOnly()
	if err != nil {
		sendJSONError(w, "Could not load database to create the new user", "UNABLE_TO_LOAD_DATABASE", http.StatusInternalServerError)
		return
	}

	// Verify if the user already exists
	value, _ := checkUserExists(signUpData.Username, signUpData.Email, db)
	if value == 1 {
		sendJSONError(w, "Email already used", "EMAIL_USED", http.StatusConflict)
		return
	} else if value == 2 {
		sendJSONError(w, "Username already exists", "USERNAME_EXISTS", http.StatusConflict)
		return
	}

	// Save the user in the database, need to enter x2 password (UI design for the local app)
	newUser, err := mf.NewUser(signUpData.Username, signUpData.Email, signUpData.Password, signUpData.Password, db)
	if err != nil {
		sendJSONError(w, "Could not create user", "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}
	errS := mdb.SaveDB(db)
	if errS != nil {
		sendJSONError(w, "Could not save the new user", "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	// User Id is the number of current registered users (there is possibility yet to delete a user, so this should work for now)
	credToken, err := generateJWT(newUser.ID)
	if err != nil {
		sendJSONError(w, "Could not generate credToken", "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	log.Println("API : creation of a session")
	createSession(w, r, credToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"cred_token": credToken})
}

// CheckPersistenceHandler process the request to verify if the current user has been connected within a certain duration (1h)
func checkPersistenceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API : Received request to check the persistence")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get userID from the header
	userID := r.Header.Get("User-ID")

	// persistence == false
	if userID == "" {
		log.Println("API : connexion is fresh")
		response := map[string]any{
			"authenticated": false,
			"database":      nil,
			"user_id":       uuid.UUID{},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		// persistence == true
		id, err := uuid.Parse(userID)
		log.Println("API : persistence is available")
		if err != nil {
			log.Println("API : Failed to parse user ID")
			sendJSONError(w, "Failed to parse user ID", "PARSE_USERID", http.StatusUnauthorized)
			return
		}

		mdb.SetUserIDOfSession(id)
		db, err := mdb.LoadDB()
		if err != nil {
			log.Println("API : Failed to connect to database")
			sendJSONError(w, "Failed to connect to database", "DATABASE_CONNEXION", http.StatusInternalServerError)
			return
		}

		response := map[string]any{
			"authenticated": true,
			"database":      db,
			"user_id":       id,
		}

		log.Println("API : persistence is done")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

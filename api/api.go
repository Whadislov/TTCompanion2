package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mdb "github.com/Whadislov/TTCompanion2/internal/my_db"
	"github.com/joho/godotenv"
)

func RegisterRoutes(mux *http.ServeMux) {

	// Set env variables
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Cannot load variables from .env")
	}

	// get secrets
	SetJWTSecretKey(os.Getenv("JWT_SECRET_KEY"))
	mdb.SetPsqlInfo(os.Getenv("WEB_DB_LINK"))
	mdb.SetDBName(os.Getenv("DB_NAME"))

	// set the endpoints
	mux.Handle("/api/healthz", CorsMiddleware(http.HandlerFunc(IsApiReadyHandler)))
	mux.Handle("/api/load-database", CorsMiddleware(authMiddleware(http.HandlerFunc(loadUserDatabaseHandler))))
	mux.Handle("/api/save-database", CorsMiddleware(authMiddleware(http.HandlerFunc(saveDatabaseHandler))))
	mux.Handle("/api/login", CorsMiddleware(http.HandlerFunc(LoginHandler)))
	mux.Handle("/api/logout", CorsMiddleware(http.HandlerFunc(LogoutHandler)))
	mux.Handle("/api/signup", CorsMiddleware(http.HandlerFunc(SignUpHandler)))
	mux.Handle("/api/check-persistence", persisMiddleware(http.HandlerFunc(checkPersistenceHandler)))

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to TT Companion's API")
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to TT Companion")
	})
}

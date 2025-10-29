package myfrontend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// SignUp requests a new user creation, if everything is fine, the database of the new user and the credential token are returned
func SignUp(username string, password string, email string) (*mt.Database, string, error) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	var response struct {
		CredToken string `json:"cred_token"`
	}

	var db *mt.Database
	data.Username = username
	data.Password = password
	data.Email = email

	jsonData, err := json.Marshal(data)
	if err != nil {
		return db, "", err
	}

	resp, err := http.Post(apiURL+"signup", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return db, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse struct {
			Error string `json:"error"`
			Code  string `json:"code"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, "", fmt.Errorf("error decoding server response: %w", err)
		}

		if errorResponse.Code == "USERNAME_EXISTS" {
			return nil, "", fmt.Errorf("username is already taken")
		} else if errorResponse.Code == "EMAIL_USED" {
			return nil, "", fmt.Errorf("email is already in use")
		} else if errorResponse.Code == "UNABLE_TO_LOAD_DATABASE" {
			return nil, "", fmt.Errorf("unable to load the database")
		} else if errorResponse.Code == "INVALID_REQUEST" {
			return nil, "", fmt.Errorf("invalid request")
		} else {
			return nil, "", fmt.Errorf("internal error: %s", errorResponse.Error)
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return db, "", fmt.Errorf("error decoding JSON: %w", err)
	} else {
		log.Println("Succeed to sign user %w in", username)
		db, err := LoadDB(response.CredToken)
		if err != nil {
			return db, "", fmt.Errorf("failed to load database: %w", err)
		}
		return db, response.CredToken, nil
	}
}

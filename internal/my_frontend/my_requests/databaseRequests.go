package myfrontend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// LoadDB loads the database.
func LoadDB(credToken string) (*mt.Database, error) {
	var golangDB *mt.Database

	req, err := http.NewRequest("GET", apiURL+"load-database", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+credToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch database: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-OK status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&golangDB)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	log.Println("Database loaded successfully")
	return golangDB, nil
}

// SaveDB saves the database.
func SaveDB(credToken string, golangDB *mt.Database) error {

	dataToSave, err := json.Marshal(golangDB)
	if err != nil {
		return fmt.Errorf("failed to marshal database: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL+"save-database", bytes.NewBuffer(dataToSave))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add the token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+credToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-OK status: %d", resp.StatusCode)
	}

	log.Println("Database saved successfully")
	return nil
}

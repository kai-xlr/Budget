package storage

import (
	"encoding/json"
	"os"

	"github.com/kai-xlr/Budget/internal/models"
)

const dataFile = "budget_data.json"

func SaveSession(s models.Session) error {
	sessions := LoadSessions()
	sessions = append(sessions, s)
	data, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func LoadSessions() []models.Session {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return []models.Session{}
	}
	var sessions []models.Session
	err = json.Unmarshal(data, &sessions)
	if err != nil {
		return []models.Session{}
	}
	return sessions
}

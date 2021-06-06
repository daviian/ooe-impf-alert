package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type slot struct {
	StartDate time.Time
}

func getCovidVaccinationSlots(authority int) ([]slot, error) {
	url := fmt.Sprintf("https://e-gov.ooe.gv.at/at.gv.ooe.cip/services/api/covid/slots?page=1&size=1000&orgUnitId=%d&birthdate=1990-01-01", authority)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get slots failed: %w", err)
	}
	defer response.Body.Close()

	var jsonSlots []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&jsonSlots)

	slots, err := parseSlots(jsonSlots)
	if err != nil {
		return nil, err
	}

	return slots, err
}

// startDate of slots doesn't follow the ISO 8601 definition so we need to parse the dates by ourselves.
func parseSlots(jsonSlots []map[string]interface{}) ([]slot, error) {
	slots := make([]slot, len(jsonSlots))
	for i, jsonSlot := range jsonSlots {
		startDate, err := time.ParseInLocation("2006-01-02T15:04:05", jsonSlot["startDate"].(string), time.Local)
		if err != nil {
			return nil, fmt.Errorf("parse time failed: %w", err)
		}
		slots[i] = slot{startDate}
	}
	return slots, nil
}

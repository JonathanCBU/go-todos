package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
)

func isValidHeader(header string) bool {
	var validHeaders = []string{
		"id",
		"title",
		"description",
		"status",
		"created_timestamp",
		"updated_timestamp",
		"priority",
	}
	return slices.Contains(validHeaders, header)
}

func ReadFile(filePath string) ([]string, [][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV: %w", err)
	}
	if len(records) == 0 {
		return nil, nil, fmt.Errorf("CSV file is empty")
	}
	headers := records[0]
	for _, header := range headers {
		if !isValidHeader(header) {
			return nil, nil, fmt.Errorf("%s is not a valid header", header)
		}

	}
	var body [][]string
	if len(records) > 1 {
		body = records[1:]
	}

	return headers, body, nil
}

func HeaderMap(headers []string) map[int]string {
	var returnMap map[int]string
	for idx, h := range headers {
		returnMap[idx] = h
	}
	return returnMap
}

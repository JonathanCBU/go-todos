package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/JonathanCBU/gotodos/internal/config"
)

func ReadFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}
	headers := records[0]
	for _, header := range headers {
		err := config.ValidateHeader(header)
		if err != nil {
			return nil, err
		}

	}
	var body [][]string
	if len(records) > 1 {
		body = records[1:]
	}

	return body, nil
}

package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/JonathanCBU/gotodos/internal/models"
)

func WriteRecord(list models.List, record models.Record) error {
	fileName := list.FileName

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.Write(record.Writable())

	if err != nil {
		return fmt.Errorf("failed to write record to file: %w", err)
	}

	writer.Flush()
	err = writer.Error()
	if err != nil {
		return fmt.Errorf("failure to flush file: %w", err)
	}

	return nil
}

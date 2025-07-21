package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JonathanCBU/gotodos/internal/config"
)

type Record struct {
	Id                int
	Title             string
	Description       string
	Status            string
	Created_timestamp time.Time
	Updated_timestamp time.Time
	Priority          int
}

func NewRecord(row []string) (*Record, error) {

	if len(config.GetHeaders()) != len(row) {
		return nil, fmt.Errorf("header count (%d) doesn't match row count (%d)", len(config.GetHeaders()), len(row))
	}

	record := &Record{}

	data := make(map[string]string)
	for i, header := range config.GetHeaders() {
		data[header] = row[i]
	}

	id, err := strconv.Atoi(data["id"])
	if err != nil {
		return nil, fmt.Errorf("failure validating id: %w", err)
	}
	record.Id = id

	record.Title = data["title"]
	record.Description = data["description"]

	status := data["status"]
	err = config.ValidateStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failure validating status: %w", err)
	}
	record.Status = status

	created, err := time.Parse(config.GetTimestampFormat(), data["created_timestamp"])
	if err != nil {
		return nil, fmt.Errorf("failure vailidating created timestamp: %w", err)
	}
	record.Created_timestamp = created
	record.Updated_timestamp = created

	return record, nil
}

func (r Record) Writable() []string {
	return []string{
		strconv.Itoa(r.Id),
		r.Title,
		r.Description,
		string(r.Status),
		r.Created_timestamp.Format(config.GetTimestampFormat()),
		r.Updated_timestamp.Format(config.GetTimestampFormat()),
		strconv.Itoa(r.Priority),
	}
}

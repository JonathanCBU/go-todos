package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Status string

const (
	StatusNotStarted Status = "NOT_STARTED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusDone, StatusInProgress, StatusNotStarted:
		return true
	default:
		return false
	}
}

func ParseStatus(s string) (Status, error) {
	status := Status(strings.ToUpper(s))
	if !status.IsValid() {
		return "", fmt.Errorf("invalid status: %s", s)
	}
	return status, nil
}

type Record struct {
	Id                int
	Title             string
	Description       string
	Status            Status
	Created_timestamp time.Time
	Updated_timestamp time.Time
	Priority          int
}

func NewRecord(headers []string, row []string) (*Record, error) {
	if len(headers) != len(row) {
		return nil, fmt.Errorf("header count (%d) doesn't match row count (%d)", len(headers), len(row))
	}

	record := &Record{}

	data := make(map[string]string)
	for i, header := range headers {
		data[header] = row[i]
	}

	id, err := strconv.Atoi(data["id"])
	if err != nil {
		return nil, fmt.Errorf("failure validating id: %w", err)
	}
	record.Id = id

	record.Title = data["title"]
	record.Description = data["description"]

	status, err := ParseStatus(data["status"])
	if err != nil {
		return nil, fmt.Errorf("failure validating status: %w", err)
	}
	record.Status = status

	created, err := time.Parse(time.DateTime, data["created_timestamp"])
	if err != nil {
		return nil, fmt.Errorf("failure vailidating created timestamp: %w", err)
	}
	record.Created_timestamp = created
	record.Updated_timestamp = created

	return record, nil
}

func (r Record) Format() string {
	record := fmt.Sprintf(
		"%d, %s, %s, %s, %s, %s, %d",
		r.Id,
		r.Title,
		r.Description,
		r.Status,
		r.Created_timestamp.String(),
		r.Updated_timestamp.String(),
		r.Priority,
	)
	return record
}

func (r *Record) SetStatus(s string) (Status, error) {
	status, err := ParseStatus(s)
	if err != nil {
		return "", err
	}
	return status, err
}

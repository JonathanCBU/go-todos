package models

import (
	"errors"
	"fmt"
	"time"
)

type Status string

const (
	StatusNotStarted Status = "NOT_STARTED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
)

func (s Status) IsValid() (bool, error) {
	switch s {
	case StatusDone, StatusInProgress, StatusNotStarted:
		return true, nil
	default:
		errorMsg := fmt.Sprintf("%s is not a valid status!", s)
		return false, errors.New(errorMsg)
	}
}

type Record struct {
	id                uint8
	title             string
	descritption      string
	status            Status
	created_timestamp time.Time
	updated_timestamp time.Time
	priority          uint8
}

func (r Record) Format() string {
	record := fmt.Sprintf(
		"%d, %s, %s, %s, %s, %s, %d",
		r.id,
		r.title,
		r.descritption,
		r.status,
		r.created_timestamp.String(),
		r.updated_timestamp.String(),
		r.priority,
	)
	return record
}

func (r *Record) SetStatus(s string) (Status, error) {
	status := Status(s)

	isValid, err := status.IsValid()
	if !isValid && err == nil {
		r.status = status
	}

	return status, err
}

func (r *Record) SetCreatedDateTime(timestamp string) (time.Time, error) {
	time, err := time.Parse(time.DateTime, timestamp)
	if err == nil {
		r.created_timestamp = time
	}

	return time, err
}

func (r *Record) SetUpdatedDateTime(timestamp string) (time.Time, error) {
	time, err := time.Parse(time.DateTime, timestamp)
	if err == nil {
		r.updated_timestamp = time
	}

	return time, err
}

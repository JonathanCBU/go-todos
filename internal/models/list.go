package models

import (
	"errors"
	"fmt"
)

type List struct {
	FileName string
	Records  map[int]*Record
}

func (l *List) AddTodo(todo *Record) (int, error) {

	_, todoInList := l.Records[todo.Id]
	if todoInList {
		errorMsg := errors.New("duplicate todo ID")
		return len(l.Records), errorMsg
	}

	l.Records[todo.Id] = todo

	return len(l.Records), nil
}

func NewList(fileName string, table [][]string) (*List, error) {
	list := &List{
		FileName: fileName,
		Records:  make(map[int]*Record),
	}
	var record *Record
	var err error

	for _, row := range table {
		record, err = NewRecord(row)
		if err != nil {
			return nil, fmt.Errorf("failure generating record: %w", err)
		}
		_, err = list.AddTodo(record)
		if err != nil {
			return nil, fmt.Errorf("failure mapping record %d to list: %w", record.Id, err)
		}
	}

	return list, nil
}

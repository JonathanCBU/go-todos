package main

import (
	"fmt"
	"log"

	"github.com/JonathanCBU/gotodos/internal/csv"
	"github.com/JonathanCBU/gotodos/internal/models"
)

func main() {
	header, body, err := csv.ReadFile("/Users/jonathancook/dev/personal/go-todos/example.csv")

	if err != nil {
		log.Fatal(err)
	}

	list, err := models.NewList(
		"/Users/jonathancook/dev/personal/go-todos/example.csv",
		header,
		body,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list.FileName)
	fmt.Println(list.Records)

	for idx, row := range list.Records {
		fmt.Println(idx, row.Description)
	}

	r1, err := models.NewRecord(
		list.Headers,
		[]string{
			"666",
			"This is new",
			"Example description test",
			"DONE",
			"2006-01-02 15:04:05",
			"2006-01-06 15:04:05",
			"1",
		},
	)

	fmt.Println(r1.Writable())

	if err != nil {
		log.Fatal(err)
	}

	err = csv.WriteRecord(*list, *r1)

	if err != nil {
		log.Fatal(err)
	}

}

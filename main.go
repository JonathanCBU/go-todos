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
}

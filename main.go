package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JonathanCBU/gotodos/internal/config"
	"github.com/JonathanCBU/gotodos/internal/csv"
	"github.com/JonathanCBU/gotodos/internal/models"
)

func main() {
	err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}

	body, err := csv.ReadFile("/Users/jonathancook/dev/personal/go-todos/example.csv")

	if err != nil {
		log.Fatal(err)
	}

	list, err := models.NewList(
		"/Users/jonathancook/dev/personal/go-todos/example.csv",
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

	// r1, err := models.NewRecord(
	// 	list.Headers,
	// 	[]string{
	// 		"666",
	// 		"This is new",
	// 		"Example description test",
	// 		"DONE",
	// 		"2006-01-02 15:04:05",
	// 		"2006-01-06 15:04:05",
	// 		"1",
	// 	},
	// )

	// fmt.Println(r1.Writable())

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = csv.WriteRecord(*list, *r1)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var i int
	// fmt.Print("Type a number: ")
	// fmt.Scan(&i)
	// fmt.Println("Your number is:", i)
	fmt.Println(time.DateTime)
}

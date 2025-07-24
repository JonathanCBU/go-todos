package cmd

import (
	"os"

	"log"

	"github.com/JonathanCBU/gotodos/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godos",
	Short: "Basic todo app in Go",
}

func Execute() {
	// err := config.Load("config.json")
	err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}

package cmd

import (
	"fmt"

	"github.com/JonathanCBU/gotodos/internal/config"
	"github.com/spf13/cobra"
)

var listConfig = &cobra.Command{
	Use:   "config",
	Short: "List configs for the todo app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetTimestampFormat())
	},
}

func init() {
	rootCmd.AddCommand(listConfig)
}

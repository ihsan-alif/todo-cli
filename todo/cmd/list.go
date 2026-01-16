/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ihsan-alif/todo-cli/todo/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all to-do items",
	Run: func(cmd *cobra.Command, args []string) {
		// check if there are no todos
		if len(internal.Todos) == 0 {
			fmt.Println("No to-do items found")
			return
		}

		// print all todos
		for _, t := range internal.Todos {
			status := "❌"
			if t.Done {
				status = "✅"
			}
			fmt.Printf("%d. %s [%s]\n", t.ID, t.Text, status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

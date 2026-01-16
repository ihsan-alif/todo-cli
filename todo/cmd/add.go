/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/ihsan-alif/todo-cli/todo/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add new to-do item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		now := time.Now()
		// create new to-do item
		todo := internal.Todo{
			ID:        len(internal.Todos) + 1,
			Text:      args[0],
			Done:      false,
			CreatedAt: now,
			UpdatedAt: now,
		}

		// add to-do item to the list
		internal.Todos = append(internal.Todos, todo)
		// save to-do list to file
		internal.Savetodos()
		fmt.Printf("Task added successfully (ID: %v)\n", todo.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}

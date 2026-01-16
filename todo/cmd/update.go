/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ihsan-alif/todo-cli/todo/internal"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [new text]",
	Short: "Update to-do item",
	Run: func(cmd *cobra.Command, args []string) {
		// get id from args
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID must be a number")
			return
		}
		// get new text from args
		newText := args[1]

		// find and update to-do item
		for i, t := range internal.Todos {
			if t.ID == id {
				internal.Todos[i].Text = newText
				internal.Savetodos()
				fmt.Println("To-do item successfully updated")
				return
			}
		}
		fmt.Println("To-do item is not found")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}

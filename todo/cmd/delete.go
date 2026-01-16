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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a to-do item",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID must be a number")
			return
		}

		for i, t := range internal.Todos {
			if t.ID == id {
				internal.Todos = append(internal.Todos[:i], internal.Todos[i+1:]...)
				internal.Savetodos()
				fmt.Println("to-do item deleted")
				return
			}
		}
		fmt.Println("to-do item not found")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}

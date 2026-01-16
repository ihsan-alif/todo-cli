/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ihsan-alif/todo-cli/todo/internal"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a to-do item as done",
	Run: func(cmd *cobra.Command, args []string) {
		// get id from args
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID must be a number")
			return
		}

		// find and mark to-do item as done
		for i, t := range internal.Todos {
			if t.ID == id {
				internal.Todos[i].Done = true
				internal.Todos[i].UpdatedAt = time.Now()
				internal.Savetodos()
				fmt.Printf("%v marked as done\n", t.Text)
				return
			}
		}
		fmt.Printf("to-do item with ID %d not found\n", id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}

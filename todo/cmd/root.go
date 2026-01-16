/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ihsan-alif/todo-cli/todo/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "simple CLI to-do application with cobra",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		internal.LoadTodos()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}

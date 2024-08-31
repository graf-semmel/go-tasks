/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Add(args[0])
		ListPending()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	// addCmd.Flags().StringVarP(&desc, "description", "d", "", "Task description")
	// addCmd.Flags().BoolP("done", "D", false, "Is the task done?")
	// addCmd.Flags().StringVarP(&due, "due", "t", "", "Due date")
	// addCmd.MarkFlagRequired("description")
}

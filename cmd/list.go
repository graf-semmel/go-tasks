/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var all bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			ListAll()
		} else {
			ListPending()
		}
	},
}

func ListAll() {
	w := newTabWriter()
	defer w.Flush()

	fmt.Println("=== All tasks ===")
	fmt.Fprintln(w, "ID\tDone\tTask\tCreated\t")
	tasks := GetAll()
	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t\n", task.ID, toCheckmark(task.Done), task.Description, toTimeDiff(task.Created))
	}
}

func ListPending() {
	tasks := GetPending()

	if len(tasks) == 0 {
		fmt.Println("Good job! All tasks are done.")
		return
	}

	w := newTabWriter()
	defer w.Flush()

	fmt.Println("=== TODO ===")
	fmt.Fprintln(w, "ID\tTask\tCreated\t")
	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t\n", task.ID, task.Description, toTimeDiff(task.Created))
	}
}

func ListSingle(task Task) {
	w := newTabWriter()
	defer w.Flush()

	fmt.Printf("=== Task with ID %d ===\n", task.ID)
	fmt.Fprintln(w, "ID\tDone\tTask\tCreated\t")
	fmt.Fprintf(w, "%d\t%s\t%s\t%s\t\n", task.ID, toCheckmark(task.Done), task.Description, toTimeDiff(task.Created))
}

func toTimeDiff(stringRFC3339 string) string {
	time, _ := time.Parse(time.RFC3339, stringRFC3339)
	return timediff.TimeDiff(time)
}

func toCheckmark(done bool) string {
	if done {
		return "✓"
	}
	return "✗"
}

func newTabWriter() *tabwriter.Writer {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	return w
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all tasks")
}

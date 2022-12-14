package cmd

import (
	"fmt"
	"strings"
	"task/db"
	"task/helpers"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		helpers.Must(err, fmt.Sprintf("There was an error adding your task: '%s'", task))
		fmt.Printf("Added '%s' to your task list\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

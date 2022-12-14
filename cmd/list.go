package cmd

import (
	"fmt"
	"task/db"
	"task/helpers"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your current tasks",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command.
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadAllTasks()
		helpers.Must(err, "Error listing all your tasks")
		if len(tasks) == 0 {
			fmt.Println("You have no tasks, congrats you hard worker :)")
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

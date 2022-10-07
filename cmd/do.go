/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"task/db"
	"task/helpers"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks_index []int
		for _, value := range args {
			i, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Could not do task")
				fmt.Println("error: ", err.Error())
			} else {
				tasks_index = append(tasks_index, i)
			}
		}
		tasks, err := db.ReadAllTasks()
		helpers.Must(err, "Error reading all your tasks")
		var task_ids []int
		for _, index := range tasks_index {
			i := index - 1
			if i >= 0 && i < len(tasks) {
				task_ids = append(task_ids, tasks[i].Key)
				fmt.Printf("Marking '%s' off your list\n", tasks[i].Value)
			} else {
				fmt.Printf("Task '%d' does not exist \n", index)
			}
		}
		for _, task_id := range task_ids {
			err := db.DeleteTask(task_id)
			if err != nil {
				fmt.Println("Could not delete task with id:", task_id)
				fmt.Println("error: ", err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

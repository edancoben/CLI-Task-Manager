/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"task/db"

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
		fmt.Println("do called")
		var task_ids []int
		for _, value := range args {
			i, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Could not do task with id:", value)
				fmt.Println("error: ", err)
			} else {
				task_ids = append(task_ids, i)
			}
		}
		fmt.Println(task_ids)
		for _, task_id := range task_ids {
			err := db.DeleteTask(task_id)
			if err != nil {
				fmt.Println("Could delete task with id:", task_id)
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

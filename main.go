/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"path/filepath"
	"task/cmd"
	"task/db"
	"task/helpers"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home_dir, dir_err := homedir.Dir()
	helpers.Must(dir_err, "error finding your home directory")
	db_dir := filepath.Join(home_dir, ".tasks.db")
	helpers.Must(db.Init(db_dir), "error starting db")
	cmd.Execute()
}

# CLI-Task-Manager

A simple CLI task manager built in GO with the cobra cli framework and `bolt db` backend


Usage:
  task [command]

Available Commands:
  add         Adds a task to your task list
  completion  Generate the autocompletion script for the specified shell
  do          Crosses tasks off your list specified by their id
  help        Help about any command
  list        Lists all your current tasks

Flags:
  -h, --help     help for task
  -t, --toggle   Help message for toggle
# CLI TASK MANAGER
In this exercise we are creating command line tool for todo tasks.
```
You can perform task like:
1. add task to your todos
2. do the task from todos
3. list the task in your todos
```
```
$ task
With task command you can add task to your todo list do the task and view all the task

Usage:
  task [command]

Available Commands:
  add         add command is used to add task into todo
  do          do command is used for doing the task from todo list
  help        Help about any command
  list        list command will list all the task in the todo

Flags:
  -h, --help   help for task

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal"

$ task add clean dishes
Added "clean dishes"

$ task list
1. review talk proposal
2. some task description

$ task do 1

$ task list
You have the following tasks:
1. some task description
```

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BabaykaBo/todo-app"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add new todo")
	complete := flag.Int("add", 0, "add new todo")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "Invalid command")
		os.Exit(0)
	}
}

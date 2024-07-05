package main

import (
	"fmt"
	"os"
	"strings"
)

const dbPath = ".todo.pb"

// list lists all tasks by calling the gRpc server
func list() error {
	return fmt.Errorf("not implemented")
}

// add adds a task by using a gRpc client to call the gRpc server
func add(description string) error {
	return fmt.Errorf("not implemented")
}

// markDone marks the task at index indexStr done by calling the gRpc server
func markDone(indexStr string) error {
	return fmt.Errorf("not implemented")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: todo <command> [arguments]\n")
		return
	}
	cmd := os.Args[1]
	var err error
	switch cmd {
	case "add":
		err = add(strings.Join(os.Args[2:], " "))
	case "list":
		err = list()
	case "done":
		if len(os.Args) != 3 {
			err = fmt.Errorf("usage: todo done <index>")
			break
		}
		err = markDone(os.Args[2])
	default:
		err = fmt.Errorf("unknown command %q", cmd)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

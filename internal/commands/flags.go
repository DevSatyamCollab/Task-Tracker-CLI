package commands

import (
	"flag"
	"fmt"
	"os"
)

const (
	NoIDSelected = -1
)

type CmdFlag struct {
	Add          string
	Description  string
	List         bool
	Todo         bool
	Progress     bool
	Done         bool
	MarkDone     int
	MarkProgress int
	MarkTodo     int
	Update       int
	Delete       int
}

func ParseFlags(args []string) (*CmdFlag, error) {
	flags := &CmdFlag{}
	fs := flag.NewFlagSet("todo-list", flag.ContinueOnError)

	// set custom usage
	fs.Usage = func() {
		usage := `Usage: todo-list [options]

Examples:
  todo-list -add "Buy a Coffee"
  todo-list -update 1 -desc "Go for Running"
  todo-list -delete 1
  todo-list -list
  todo-list -list -todo 
  todo-list -list -progress
  todo-list -list -done
  todo-list -markprogress 1
  todo-list -markdone 1
  todo-list -marktodo 1

Options:`
		fmt.Fprintln(os.Stderr, usage)
		fs.PrintDefaults()
	}

	// define flags
	fs.StringVar(&flags.Add, "add", "", "Add a new task")
	fs.IntVar(&flags.Delete, "del", NoIDSelected, "Delete task by ID")
	fs.IntVar(&flags.Update, "update", NoIDSelected, "Update task by ID")
	fs.StringVar(&flags.Description, "desc", "", "description of task")

	fs.IntVar(&flags.MarkDone, "markdone", NoIDSelected, "mark as done by ID")
	fs.IntVar(&flags.MarkProgress, "markprogress", NoIDSelected, "mark as progress by ID")
	fs.IntVar(&flags.MarkTodo, "marktodo", NoIDSelected, "mark as todo by ID")
	fs.BoolVar(&flags.List, "list", false, "List all Task")
	fs.BoolVar(&flags.Done, "done", false, "list Done Task")
	fs.BoolVar(&flags.Progress, "progress", false, "list Progress Task")
	fs.BoolVar(&flags.Todo, "todo", false, "list pending Task")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return flags, nil
}

// DetermineCommand determines which command to execute based on flags
func (f *CmdFlag) DetermineCommand() (string, error) {
	commandCount := 0
	var command string

	if f.Add != "" {
		commandCount++
		command = "add"
	}
	if f.List {
		commandCount++
		command = "list"
	}
	if f.MarkDone != NoIDSelected {
		commandCount++
		command = "markDone"
	}

	if f.MarkProgress != NoIDSelected {
		commandCount++
		command = "markProgress"
	}

	if f.Delete != NoIDSelected {
		commandCount++
		command = "delete"
	}
	if f.Update != NoIDSelected {
		commandCount++
		command = "update"
	}
	if f.MarkTodo != NoIDSelected {
		commandCount++
		command = "marktodo"
	}

	if commandCount == 0 {
		return "", fmt.Errorf("no command specified. Use -h for help")
	}
	if commandCount > 1 {
		return "", fmt.Errorf("multiple commands specified. Use only one command at a time")
	}

	return command, nil
}

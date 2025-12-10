package app

import (
	"todo-list/internal/commands"
	"todo-list/internal/presenter"
	"todo-list/internal/service"
)

type App struct {
	service   *service.TaskService
	presenter *presenter.ConsolePresenter
}

func NewApp(s *service.TaskService, p *presenter.ConsolePresenter) *App {
	return &App{service: s, presenter: p}
}

func (a *App) createCommand(commandName string, flag *commands.CmdFlag) commands.ICommand {
	switch commandName {
	case "add":
		return commands.NewAddCommand(a.service, a.presenter, flag.Add)
	case "update":
		return commands.NewUpdateCommand(a.service, a.presenter, flag.Update, flag.Description)
	case "delete":
		return commands.NewDeleteCommand(a.service, a.presenter, flag.Delete)
	case "list":
		return commands.NewListCommand(a.service, a.presenter, flag.Todo, flag.Progress, flag.Done)
	case "markDone":
		return commands.NewDoneCommand(a.service, a.presenter, flag.MarkDone)
	case "markProgress":
		return commands.NewProgressCommand(a.service, a.presenter, flag.MarkProgress)
	case "marktodo":
		return commands.NewTodoCommand(a.service, a.presenter, flag.MarkTodo)
	default:
		return nil
	}
}

// Run the application
func (a *App) Run(args []string) error {
	// parse flags
	flags, err := commands.ParseFlags(args)
	if err != nil {
		return err
	}

	// determine command
	commandName, err := flags.DetermineCommand()
	if err != nil {
		return err
	}

	// create command
	command := a.createCommand(commandName, flags)

	// validate
	if err := command.Validate(); err != nil {
		return err
	}

	return command.Execute()
}

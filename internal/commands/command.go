package commands

import (
	"fmt"
	"todo-list/internal/core"
	"todo-list/internal/presenter"
	"todo-list/internal/service"
)

type ICommand interface {
	Validate() error
	Execute() error
}

// -----------------------
// Base Command
// -----------------------
type Command struct {
	service   *service.TaskService
	presenter *presenter.ConsolePresenter
}

func NewCommand(s *service.TaskService, p *presenter.ConsolePresenter) *Command {
	return &Command{service: s, presenter: p}
}

// -----------------------
// Add Command
// -----------------------
type AddCommand struct {
	Command
	description string
}

func NewAddCommand(s *service.TaskService, p *presenter.ConsolePresenter, desc string) *AddCommand {
	return &AddCommand{
		Command:     *NewCommand(s, p),
		description: desc,
	}
}

func (ac *AddCommand) Validate() error {
	// validate description
	if err := core.ValidateDescription(ac.description); err != nil {
		return err
	}

	return nil
}

func (ac *AddCommand) Execute() error {
	if err := ac.service.AddTask(ac.description); err != nil {
		return err
	}

	ac.presenter.ShowSuccess("Task added successfully")

	return nil
}

// -----------------------
// Update Command
// -----------------------
type UpdateCommand struct {
	Command
	id          int
	description string
}

func NewUpdateCommand(s *service.TaskService, p *presenter.ConsolePresenter, id int, desc string) *UpdateCommand {
	return &UpdateCommand{
		Command:     *NewCommand(s, p),
		id:          id,
		description: desc,
	}
}

func (uc *UpdateCommand) Validate() error {
	// validate id
	if err := core.ValidateID(uc.id); err != nil {
		return err
	}

	// validate description
	if err := core.ValidateDescription(uc.description); err != nil {
		return err
	}

	return nil
}

func (uc *UpdateCommand) Execute() error {
	if err := uc.service.UpdateTask(uc.id, uc.description); err != nil {
		return err
	}

	uc.presenter.ShowSuccess(fmt.Sprintf("Task %d updated successfully", uc.id))

	return nil
}

// -----------------------
// Delete Command
// -----------------------
type DeleteCommand struct {
	Command
	id int
}

func NewDeleteCommand(s *service.TaskService, p *presenter.ConsolePresenter, id int) *DeleteCommand {
	return &DeleteCommand{
		Command: *NewCommand(s, p),
		id:      id,
	}
}

func (dc *DeleteCommand) Validate() error {
	// validate id
	if err := core.ValidateID(dc.id); err != nil {
		return err
	}

	return nil
}

func (dc *DeleteCommand) Execute() error {
	if err := dc.service.DeleteTask(dc.id); err != nil {
		return err
	}

	dc.presenter.ShowSuccess(fmt.Sprintf("Task %d deleted successfully", dc.id))

	return nil
}

// -----------------------
// List Command
// -----------------------
type ListCommand struct {
	Command
	todo, progress, done bool
}

func NewListCommand(s *service.TaskService, p *presenter.ConsolePresenter, todo, progress, done bool) *ListCommand {
	return &ListCommand{
		Command:  *NewCommand(s, p),
		todo:     todo,
		progress: progress,
		done:     done,
	}
}

func (lc *ListCommand) Validate() error {

	return nil
}

func (lc *ListCommand) Execute() error {
	switch {
	case lc.todo:
		lc.presenter.ViewPendingList()
	case lc.progress:
		lc.presenter.ViewProgessList()
	case lc.done:
		lc.presenter.ViewCompletedList()
	default:
		lc.presenter.ViewAll()
	}

	return nil
}

// -----------------------
// Done Command
// -----------------------
type DoneCommand struct {
	Command
	id int
}

func NewDoneCommand(s *service.TaskService, p *presenter.ConsolePresenter, id int) *DoneCommand {
	return &DoneCommand{
		Command: *NewCommand(s, p),
		id:      id,
	}
}

func (dc *DoneCommand) Validate() error {
	if err := core.ValidateID(dc.id); err != nil {
		return err
	}

	return nil
}

func (dc *DoneCommand) Execute() error {
	if err := dc.service.MarkDone(dc.id); err != nil {
		return err
	}

	dc.presenter.ShowSuccess(fmt.Sprintf("Task %d marked as done", dc.id))
	return nil
}

// -----------------------
// Progress Command
// -----------------------
type ProgressCommand struct {
	Command
	id int
}

func NewProgressCommand(s *service.TaskService, p *presenter.ConsolePresenter, id int) *ProgressCommand {
	return &ProgressCommand{
		Command: *NewCommand(s, p),
		id:      id,
	}
}

func (pc *ProgressCommand) Validate() error {
	if err := core.ValidateID(pc.id); err != nil {
		return err
	}

	return nil
}

func (pc *ProgressCommand) Execute() error {
	if err := pc.service.MarkProgress(pc.id); err != nil {
		return err
	}

	pc.presenter.ShowSuccess(fmt.Sprintf("Task %d marked as progress", pc.id))

	return nil
}

// -----------------------
// MarkTodo Command
// -----------------------
type TodoCommand struct {
	Command
	id int
}

func NewTodoCommand(s *service.TaskService, p *presenter.ConsolePresenter, id int) *TodoCommand {
	return &TodoCommand{
		Command: *NewCommand(s, p),
		id:      id,
	}
}

func (tc *TodoCommand) Validate() error {
	// validate id
	if err := core.ValidateID(tc.id); err != nil {
		return err
	}

	return nil
}

func (tc *TodoCommand) Execute() error {
	if err := tc.service.MarkTodo(tc.id); err != nil {
		return err
	}

	tc.presenter.ShowSuccess(fmt.Sprintf("Task %d marked as todo", tc.id))

	return nil
}

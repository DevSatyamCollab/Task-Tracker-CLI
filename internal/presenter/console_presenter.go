package presenter

import (
	"fmt"
	"os"
	"strconv"
	"todo-list/internal/service"

	"github.com/aquasecurity/table"
)

const (
	dateFormat string = "02-01-2006"
)

type ConsolePresenter struct {
	service *service.TaskService
}

func NewConsolePresenter(s *service.TaskService) *ConsolePresenter {
	return &ConsolePresenter{service: s}
}

// view all list
func (p *ConsolePresenter) ViewAll() {
	tasklist := p.service.GetAllTask()
	if len(tasklist) == 0 {
		fmt.Println("List is Empty")
		return
	}

	fmt.Println("All Tasks:\n----------------------------------------------------------------")
	tble := table.New(os.Stdout)
	tble.SetHeaders("ID", "Created At", "Description", "Status", "Completed At")
	var completedAt string
	for _, task := range tasklist {
		completedAt = ""
		if task.Status == "✅" {
			completedAt = task.CompletedAt.Format(dateFormat)
		}

		tble.AddRow(strconv.Itoa(task.Id), task.CreatedAt.Format(dateFormat),
			task.Description, task.Status, completedAt)
	}

	tble.Render()
}

// pending
func (p *ConsolePresenter) ViewPendingList() {
	tasklist := p.service.GetAllPendingTask()
	if len(tasklist) == 0 {
		fmt.Println("list is empty")
		return
	}

	fmt.Println("❌ PENDING TASKS:\n----------------------------------------------------------------")
	tble := table.New(os.Stdout)
	tble.SetHeaders("ID", "Created At", "Description")

	for _, task := range tasklist {
		tble.AddRow(strconv.Itoa(task.Id), task.CreatedAt.Format(dateFormat), task.Description)
	}

	tble.Render()
}

// completed
func (p *ConsolePresenter) ViewCompletedList() {

	tasklist := p.service.GetAllCompletedTask()
	if len(tasklist) == 0 {
		fmt.Println("list is empty")
		return
	}

	fmt.Println("✅ COMPLETED TASKS:\n----------------------------------------------------------------")
	tble := table.New(os.Stdout)
	tble.SetHeaders("ID", "Created At", "Description", "Completed At")

	for _, task := range p.service.GetAllCompletedTask() {
		tble.AddRow(strconv.Itoa(task.Id), task.CreatedAt.Format(dateFormat),
			task.Description, task.CompletedAt.Format(dateFormat))
	}

	tble.Render()
}

// progress
func (p *ConsolePresenter) ViewProgessList() {

	tasklist := p.service.GetAllProgressTask()
	if len(tasklist) == 0 {
		fmt.Println("list is empty")
		return
	}

	fmt.Println("⌛ PROGRESS TASKS:\n----------------------------------------------------------------")
	tble := table.New(os.Stdout)
	tble.SetHeaders("ID", "Created At", "Description")

	for _, task := range p.service.GetAllProgressTask() {
		tble.AddRow(strconv.Itoa(task.Id), task.CreatedAt.Format(dateFormat), task.Description)
	}

	tble.Render()
}

// success msg
func (p *ConsolePresenter) ShowSuccess(msg string) {
	fmt.Println(msg)
}

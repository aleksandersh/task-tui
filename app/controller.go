package app

import (
	"context"

	"github.com/aleksandersh/task-tui/app/pages/tasks"
	"github.com/aleksandersh/task-tui/app/ui"
	"github.com/aleksandersh/task-tui/domain"
	"github.com/aleksandersh/task-tui/task"
	"github.com/rivo/tview"
)

const (
	pageNameTasks = "tasks"
)

type controller struct {
	ctx       context.Context
	task      *task.Task
	taskfile  *domain.Taskfile
	config    *ui.Config
	app       *tview.Application
	pagesView *tview.Pages
}

func newController(ctx context.Context, task *task.Task, app *tview.Application, taskfile *domain.Taskfile, config *ui.Config) ui.Controller {
	pagesView := tview.NewPages()
	return &controller{
		ctx:       ctx,
		task:      task,
		taskfile:  taskfile,
		config:    config,
		app:       app,
		pagesView: pagesView,
	}
}

func (c *controller) StartUi() {
	c.app.SetRoot(c.pagesView, true)
	c.ShowTasks()
}

func (c *controller) ShowTasks() {
	c.pagesView.AddAndSwitchToPage(pageNameTasks, tasks.New(c.ctx, c.task, c.config, c, c.taskfile), true)
}

func (c *controller) Focus(view tview.Primitive) {
	c.app.SetFocus(view)
}

func (c *controller) PostDraw(f func()) {
	c.app.QueueUpdateDraw(f)
}

func (c *controller) Close() {
	c.app.Stop()
}

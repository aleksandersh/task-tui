package app

import (
	"context"

	"github.com/aleksandersh/taskfile-tui/app/pages/tasks"
	"github.com/aleksandersh/taskfile-tui/app/ui"
	"github.com/aleksandersh/taskfile-tui/domain"
	"github.com/rivo/tview"
)

const (
	pageNameTasks = "tasks"
)

type controller struct {
	ctx       context.Context
	taskfile  *domain.Taskfile
	app       *tview.Application
	pagesView *tview.Pages
}

func newController(ctx context.Context, app *tview.Application, taskfile *domain.Taskfile) ui.Controller {
	pagesView := tview.NewPages()
	return &controller{ctx: ctx, app: app, pagesView: pagesView, taskfile: taskfile}
}

func (c *controller) StartUi() {
	c.app.SetRoot(c.pagesView, true)
	c.ShowTasks()
}

func (c *controller) ShowTasks() {
	c.pagesView.AddAndSwitchToPage(pageNameTasks, tasks.New(c.ctx, c.app, c, c.taskfile), true)
}

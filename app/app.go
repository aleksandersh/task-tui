package app

import (
	"context"
	"fmt"

	"github.com/aleksandersh/task-tui/domain"
	"github.com/aleksandersh/task-tui/task"
	"github.com/rivo/tview"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Start(ctx context.Context, task *task.Task, taskfile *domain.Taskfile) error {
	app := tview.NewApplication()
	contoller := newController(ctx, task, app, taskfile)
	contoller.StartUi()

	if err := app.Run(); err != nil {
		return fmt.Errorf("error in app.Run: %w", err)
	}
	app.Stop()
	return nil
}

package app

import (
	"context"
	"fmt"

	"github.com/aleksandersh/task-tui/app/ui"
	"github.com/aleksandersh/task-tui/cli"
	"github.com/aleksandersh/task-tui/domain"
	"github.com/aleksandersh/task-tui/task"
	"github.com/rivo/tview"
)

type App struct {
	config *ui.Config
}

func New(args *cli.Args) *App {
	cfg := ui.Config{
		SecondLineEnabled: args.EnableSecondLine,
	}
	return &App{config: &cfg}
}

func (a *App) Start(ctx context.Context, task *task.Task, taskfile *domain.Taskfile) error {
	app := tview.NewApplication()
	contoller := newController(ctx, task, app, taskfile, a.config)
	contoller.StartUi()

	if err := app.Run(); err != nil {
		return fmt.Errorf("error in app.Run: %w", err)
	}
	app.Stop()
	return nil
}

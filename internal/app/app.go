package app

import (
	"context"
	"fmt"

	"github.com/aleksandersh/task-tui/internal/app/ui"
	"github.com/aleksandersh/task-tui/internal/cli"
	"github.com/aleksandersh/task-tui/internal/domain"
	"github.com/aleksandersh/task-tui/internal/task"
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

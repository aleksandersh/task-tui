package app

import (
	"context"
	"fmt"

	"github.com/aleksandersh/taskfile-tui/domain"
	"github.com/rivo/tview"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Start(ctx context.Context, taskfile *domain.Taskfile) error {
	app := tview.NewApplication()
	contoller := newController(ctx, app, taskfile)
	contoller.StartUi()

	if err := app.Run(); err != nil {
		return fmt.Errorf("error in app.Run: %w", err)
	}
	app.Stop()
	return nil
}

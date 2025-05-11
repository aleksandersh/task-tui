package cli

import (
	"context"

	"github.com/aleksandersh/task-tui/internal/data"
	"github.com/aleksandersh/task-tui/internal/task"
)

func Repeat(ctx context.Context, args *Args) error {
	cmd, err := data.LoadLatestCommand()
	if err != nil {
		return err
	}

	task.New([]string{}, cmd.Args, cmd.CliArgs).ExecuteTask(ctx, cmd.Name)
	return nil
}

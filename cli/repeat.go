package cli

import (
	"context"

	"github.com/aleksandersh/task-tui/data"
	"github.com/aleksandersh/task-tui/task"
)

func Repeat(ctx context.Context, args *Args) error {
	cmd, err := data.LoadLatestCommand()
	if err != nil {
		return err
	}

	task.New([]string{}, cmd.Args).ExecuteTask(ctx, cmd.Name)
	return nil
}

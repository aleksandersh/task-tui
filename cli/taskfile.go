package cli

import (
	"fmt"
	"path/filepath"

	"github.com/aleksandersh/task-tui/task"
)

func CreateTask(args *Args) (*task.Task, error) {
	listArgs := []string{"--list-all", "--json", "--no-status", "--sort", args.Sort}
	taskArgs := []string{}
	if args.ExitCode {
		taskArgs = append(taskArgs, "--exit-code")
	}
	if args.Global {
		listArgs = append(listArgs, "--global")
		taskArgs = append(taskArgs, "--global")
	} else if len(args.Taskfile) > 0 {
		path, err := filepath.Abs(args.Taskfile)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve file path: %w", err)
		}
		listArgs = append(listArgs, "--taskfile", path)
		taskArgs = append(taskArgs, "--taskfile", path)
	}
	return task.New(listArgs, taskArgs), nil
}

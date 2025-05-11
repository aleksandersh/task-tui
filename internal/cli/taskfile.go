package cli

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/aleksandersh/task-tui/internal/task"
)

func CreateTask(args *Args) (*task.Task, error) {
	listArgs := []string{"--list-all", "--json"}
	taskArgs := []string{}
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

	if args.Concurrency > 0 {
		taskArgs = append(taskArgs, "--concurrency", strconv.Itoa(args.Concurrency))
	}
	if len(args.Dir) > 0 {
		listArgs = append(listArgs, "--dir", args.Dir)
		taskArgs = append(taskArgs, "--dir", args.Dir)
	}
	if args.Dry {
		taskArgs = append(taskArgs, "--dry")
	}
	if args.ExitCode {
		taskArgs = append(taskArgs, "--exit-code")
	}
	if args.Force {
		taskArgs = append(taskArgs, "--force")
	}
	if args.Interval > 0 {
		taskArgs = append(taskArgs, "--interval", args.Interval.String())
	}
	if len(args.TaskSort) > 0 {
		listArgs = append(listArgs, "--sort", args.TaskSort)
	}

	if len(args.Output) > 0 {
		taskArgs = append(taskArgs, "--output", args.Output)
	}
	if len(args.OutputGroupBegin) > 0 {
		taskArgs = append(taskArgs, "--output-group-begin", args.OutputGroupBegin)
	}
	if len(args.OutputGroupEnd) > 0 {
		taskArgs = append(taskArgs, "--output-group-end", args.OutputGroupEnd)
	}
	if args.OutputGroupErrorOnly {
		taskArgs = append(taskArgs, "--output-group-error-only")
	}

	if args.Parallel {
		taskArgs = append(taskArgs, "--parallel")
	}
	if args.Silent {
		taskArgs = append(taskArgs, "--silent")
	}
	if args.AssumeYes {
		taskArgs = append(taskArgs, "--yes")
	}
	if args.Status {
		taskArgs = append(taskArgs, "--status")
	}
	if args.Insecure {
		taskArgs = append(taskArgs, "--insecure")
	}
	if args.Verbose {
		taskArgs = append(taskArgs, "--verbose")
	}
	if args.Watch {
		taskArgs = append(taskArgs, "--watch")
	}

	return task.New(listArgs, taskArgs), nil
}

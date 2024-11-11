package cli

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/aleksandersh/task-tui/task"
)

func CreateTask(args *Args) (*task.Task, error) {
	listArgs := []string{"--list-all", "--json"}
	taskArgs := []string{}
	if args.Global {
		listArgs = append(listArgs, "--global")
		taskArgs = append(taskArgs, "--global")
	} else if args.Taskfile != nil {
		path, err := filepath.Abs(*args.Taskfile)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve file path: %w", err)
		}
		listArgs = append(listArgs, "--taskfile", path)
		taskArgs = append(taskArgs, "--taskfile", path)
	}

	if args.Concurrency != nil {
		taskArgs = append(taskArgs, "--concurrency", strconv.Itoa(*args.Concurrency))
	}
	if args.Dir != nil {
		listArgs = append(listArgs, "--dir", *args.Dir)
		taskArgs = append(taskArgs, "--dir", *args.Dir)
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
	if args.Interval != nil {
		taskArgs = append(taskArgs, "--interval", *args.Interval)
	}
	if args.Sort != nil {
		listArgs = append(listArgs, "--sort", *args.Sort)
	}

	if args.Output != nil {
		taskArgs = append(taskArgs, "--output", *args.Output)
	}
	if args.OutputGroupBegin != nil {
		taskArgs = append(taskArgs, "--output-group-begin", *args.OutputGroupBegin)
	}
	if args.OutputGroupEnd != nil {
		taskArgs = append(taskArgs, "--output-group-end", *args.OutputGroupEnd)
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
	if args.Yes {
		taskArgs = append(taskArgs, "--yes")
	}
	if args.Status {
		taskArgs = append(taskArgs, "--status")
	}
	if args.Verbose {
		taskArgs = append(taskArgs, "--verbose")
	}
	if args.Watch {
		taskArgs = append(taskArgs, "--watch")
	}

	return task.New(listArgs, taskArgs), nil
}

package main

import (
	"context"
	"log"
	"os"

	"github.com/aleksandersh/task-tui/internal/app"
	"github.com/aleksandersh/task-tui/internal/cli"
)

func main() {
	args := cli.GetArgs(os.Args)

	if args.Help {
		args.PrintUsage()
		return
	}

	if args.Version {
		args.PrintVersion()
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if args.Repeat {
		if err := cli.Repeat(ctx, args); err != nil {
			log.Fatalf("failed to repeat latest command: %v", err)
		}
		return
	}

	task, err := cli.CreateTask(args)
	if err != nil {
		log.Fatalf("failed to setup taskfile: %v", err)
	}

	taskfile, err := task.LoadTaskfile(ctx)
	if err != nil {
		log.Fatalf("failed to load taskfile: %v", err)
	}

	if err := app.New(args).Start(ctx, task, taskfile); err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}

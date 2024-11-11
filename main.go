package main

import (
	"context"
	"log"

	"github.com/aleksandersh/task-tui/app"
	"github.com/aleksandersh/task-tui/cli"
)

func main() {
	args := cli.GetArgs()

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

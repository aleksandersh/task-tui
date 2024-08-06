package main

import (
	"context"
	"log"

	"github.com/aleksandersh/task-tui/app"
	"github.com/aleksandersh/task-tui/cli"
	"github.com/aleksandersh/task-tui/task"
)

func main() {
	args := cli.GetArgs()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	task := task.New(args)
	taskfile, err := task.LoadTaskfile(ctx)
	if err != nil {
		log.Fatalf("failed to load taskfile: %v", err)
	}

	if err := app.New(args).Start(ctx, task, taskfile); err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}

package main

import (
	"context"
	"log"

	"github.com/aleksandersh/taskfile-tui/app"
	"github.com/aleksandersh/taskfile-tui/cli"
	"github.com/aleksandersh/taskfile-tui/loader"
)

func main() {
	args := cli.GetArgs()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	taskfile, err := loader.New(args.Config).LoadTaskfile(ctx)
	if err != nil {
		log.Fatalf("failed to load taskfile: %v", err)
	}

	if err := app.New().Start(ctx, taskfile); err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}

package task

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

type script struct {
	ctx context.Context
	cmd *exec.Cmd
}

func newScript(ctx context.Context, cmd *exec.Cmd) *script {
	return &script{ctx: ctx, cmd: cmd}
}

func (s *script) execute() {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	startCmdAsync(ctx, s.cmd)
	go propogateSystemSignals(ctx, s.cmd)
	awaitCmdCompletion(s.cmd)
}

func startCmdAsync(ctx context.Context, cmd *exec.Cmd) {
	if err := cmd.Start(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			log.Fatalf("error in cmd.Start: %v", err)
		}
	}
}

func propogateSystemSignals(ctx context.Context, cmd *exec.Cmd) {
	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-ctx.Done():
			return
		case signal := <-signals:
			cmd.Process.Signal(signal)
		}
	}
}

func awaitCmdCompletion(cmd *exec.Cmd) {
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			log.Fatalf("error in cmd.Wait: %v", err)
		}
	}
}

package loader

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/aleksandersh/taskfile-tui/domain"
)

type TaskfileLoader struct {
	file string
}

func New(file string) *TaskfileLoader {
	return &TaskfileLoader{file: file}
}

func (l *TaskfileLoader) LoadTaskfile(ctx context.Context) (*domain.Taskfile, error) {
	var cmd *exec.Cmd
	if len(l.file) > 0 {
		cmd = createScriptCommand(ctx, "task", []string{"-c", l.file, "--list-all", "--json"})
	} else {
		cmd = createScriptCommand(ctx, "task", []string{"--list-all", "--json"})
	}

	var buffer bytes.Buffer
	cmd.Stdout = &buffer

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	startScriptAsync(cmd)
	handleSystemSignalsAsync(ctx, signals, cmd)
	awaitForScriptCompletion(cmd)

	taskfile := &domain.Taskfile{}
	err := json.Unmarshal(buffer.Bytes(), taskfile)
	return taskfile, err
}

func createScriptCommand(ctx context.Context, name string, args []string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stderr = os.Stderr
	return cmd
}

func startScriptAsync(cmd *exec.Cmd) {
	if err := cmd.Start(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			log.Fatalf("error in cmd.Start: %v", err)
		}
	}
}

func handleSystemSignalsAsync(ctx context.Context, signals chan os.Signal, cmd *exec.Cmd) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case signal := <-signals:
				cmd.Process.Signal(signal)
			}
		}
	}()
}

func awaitForScriptCompletion(cmd *exec.Cmd) {
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			log.Fatalf("error in cmd.Wait: %v", err)
		}
	}
}

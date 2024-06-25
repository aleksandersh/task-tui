package task

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/aleksandersh/task-tui/domain"
)

type Task struct {
	file string
}

func New(file string) *Task {
	return &Task{file: file}
}

func (t *Task) LoadTaskfile(ctx context.Context) (*domain.Taskfile, error) {
	script := t.newTaskfileJsonScript(ctx)

	var buffer bytes.Buffer
	script.cmd.Stdout = &buffer

	script.execute()

	taskfile := &domain.Taskfile{}
	err := json.Unmarshal(buffer.Bytes(), taskfile)
	return taskfile, err
}

func (t *Task) ExecuteTask(ctx context.Context, name string) {
	var cmd *exec.Cmd
	if len(t.file) > 0 {
		cmd = createTaskCmd(ctx, []string{"--taskfile", t.file, name})
	} else {
		cmd = createTaskCmd(ctx, []string{name})
	}
	cmd.Stdout = os.Stdout
	newScript(ctx, cmd).execute()
}

func (t *Task) newTaskfileJsonScript(ctx context.Context) *script {
	var cmd *exec.Cmd
	if len(t.file) > 0 {
		cmd = createTaskCmd(ctx, []string{"--taskfile", t.file, "--list-all", "--json", "--no-status"})
	} else {
		cmd = createTaskCmd(ctx, []string{"--list-all", "--json", "--no-status"})
	}
	return newScript(ctx, cmd)
}

func createTaskCmd(ctx context.Context, args []string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, "task", args...)
	cmd.Stderr = os.Stderr
	return cmd
}

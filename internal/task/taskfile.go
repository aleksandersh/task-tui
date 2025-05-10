package task

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/aleksandersh/task-tui/internal/data"
	"github.com/aleksandersh/task-tui/internal/domain"
)

type Task struct {
	listArgs []string
	taskArgs []string
}

func New(listArgs []string, taskArgs []string) *Task {
	return &Task{listArgs: listArgs, taskArgs: taskArgs}
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
	data.SaveLatestCommand(domain.NewCommand(name, t.taskArgs))
	var cmd *exec.Cmd
	cmd = createTaskCmd(ctx, append(t.taskArgs, name))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	newScript(ctx, cmd).execute()
}

func (t *Task) newTaskfileJsonScript(ctx context.Context) *script {
	var cmd *exec.Cmd
	cmd = createTaskCmd(ctx, t.listArgs)
	return newScript(ctx, cmd)
}

func createTaskCmd(ctx context.Context, args []string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, "task", args...)
	cmd.Stderr = os.Stderr
	return cmd
}

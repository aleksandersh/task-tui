package task

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/aleksandersh/task-tui/cli"
	"github.com/aleksandersh/task-tui/domain"
)

type Task struct {
	listArgs []string
	taskArgs []string
}

func New(args *cli.Args) *Task {
	listArgs := []string{"--list-all", "--json", "--no-status", "--sort", args.Sort}
	taskArgs := []string{}
	if args.ExitCode {
		taskArgs = append(taskArgs, "--exit-code")
	}
	if args.Global {
		listArgs = append(listArgs, "--global")
		taskArgs = append(taskArgs, "--global")
	} else if len(args.Taskfile) > 0 {
		listArgs = append(listArgs, "--taskfile", args.Taskfile)
		taskArgs = append(taskArgs, "--taskfile", args.Taskfile)
	}
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
	var cmd *exec.Cmd
	cmd = createTaskCmd(ctx, append(t.taskArgs, name))
	cmd.Stdout = os.Stdout
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

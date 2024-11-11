package cli

import (
	"runtime/debug"

	"github.com/alexflint/go-arg"
)

type Args struct {
	ExitCode         bool   `arg:"-x,--exit-code" default:"false" help:"Pass-through the exit code of the task command."`
	Global           bool   `arg:"-g,--global" default:"false" help:"Runs global Taskfile, from $HOME/Taskfile.{yml,yaml}."`
	Sort             string `arg:"--sort" default:"default" help:"Changes the order of the tasks when listed."`
	Taskfile         string `arg:"-t,--taskfile" help:"Path to Taskfile."`
	EnableSecondLine bool   `arg:"--enable-second-line" default:"false" help:"Show the description next to the task name."`
	Repeat           bool   `arg:"-r,--repeat" default:"false" help:" Try to execute last executed command."`
}

func GetArgs() *Args {
	args := &Args{}
	arg.MustParse(args)
	return args
}

func (*Args) Version() string {
	version := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		version = info.Main.Version
	}
	return "task-tui " + version
}

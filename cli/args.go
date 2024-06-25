package cli

import "github.com/alexflint/go-arg"

type Args struct {
	ExitCode bool   `arg:"-x,--exit-code" default:"false" help:"Pass-through the exit code of the task command."`
	Global   bool   `arg:"-g,--global" default:"false" help:"Runs global Taskfile, from $HOME/Taskfile.{yml,yaml}."`
	Sort     string `arg:"--sort" default:"default" help:"Changes the order of the tasks when listed."`
	Taskfile string `arg:"-t,--taskfile" help:"Path to Taskfile."`
}

func GetArgs() *Args {
	args := &Args{}
	arg.MustParse(args)
	return args
}

func (*Args) Version() string {
	return "task-tui 0.0.2"
}

package cli

import "github.com/alexflint/go-arg"

type Args struct {
	Config string `arg:"-c,--config" help:"path to a taskfile"`
}

func GetArgs() *Args {
	args := &Args{}
	arg.MustParse(args)
	return args
}

func (*Args) Version() string {
	return "task-tui 0.0.1"
}

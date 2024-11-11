package cli

import (
	"runtime/debug"

	"github.com/alexflint/go-arg"
)

type Args struct {
	Global   bool    `arg:"-g,--global" default:"false" help:"(Task) Runs global Taskfile, from $HOME/Taskfile.{yml,yaml}."`
	Taskfile *string `arg:"-t,--taskfile" help:"(Task) Path to Taskfile."`

	Concurrency *int    `arg:"-C,--concurrency" help:"(Task) Limit number tasks to run concurrently. Zero means unlimited."`
	Dir         *string `arg:"-d,--dir" help:"(Task) Sets directory of execution."`
	Dry         bool    `arg:"-n,--dry" default:"false" help:"(Task) Compiles and prints tasks in the order that they would be run, without executing them."`
	ExitCode    bool    `arg:"-x,--exit-code" default:"false" help:"(Task) Pass-through the exit code of the task command."`
	Force       bool    `arg:"-f,--force" default:"false" help:"(Task) Forces execution even when the task is up-to-date."`
	Interval    *string `arg:"-I,--interval" help:"(Task) Sets a different watch interval when using --watch, the default being 5 seconds. This string should be a valid Go Duration."`
	Sort        *string `arg:"--sort" help:"(Task) Changes the order of the tasks when listed."`

	Output               *string `arg:"-o,--output" help:"(Task) Sets output style: [interleaved/group/prefixed]."`
	OutputGroupBegin     *string `arg:"--output-group-begin" help:"(Task) Message template to print before a task's grouped output."`
	OutputGroupEnd       *string `arg:"--output-group-end" help:"(Task) Message template to print after a task's grouped output."`
	OutputGroupErrorOnly bool    `arg:"--output-group-error-only" default:"false" help:"(Task) Swallow command output on zero exit code."`

	Parallel bool `arg:"-p,--parallel" default:"false" help:"(Task) Executes tasks provided on command line in parallel."`
	Silent   bool `arg:"-s,--silent" default:"false" help:"(Task) Disables echoing."`
	Yes      bool `arg:"-y,--yes" default:"false" help:"(Task) Assume \"yes\" as answer to all prompts."`
	Status   bool `arg:"--status" default:"false" help:"(Task) Exits with non-zero exit code if any of the given tasks is not up-to-date."`
	Verbose  bool `arg:"-v,--verbose" default:"false" help:"(Task) Enables verbose mode."`
	Watch    bool `arg:"-w,--watch" default:"false" help:"(Task) Enables watch of the given task."`

	EnableSecondLine bool `arg:"--enable-second-line" default:"false" help:"Show the description next to the task name."`
	Repeat           bool `arg:"-r,--repeat" default:"false" help:"Execute last executed command."`
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

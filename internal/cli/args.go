package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/aleksandersh/task-tui/internal/build"
	"github.com/spf13/pflag"
)

type Args struct {
	EnableSecondLine bool
	Repeat           bool

	Version bool
	Help    bool

	TaskSort             string
	Status               bool
	Insecure             bool
	Watch                bool
	Verbose              bool
	Silent               bool
	AssumeYes            bool
	Parallel             bool
	Dry                  bool
	ExitCode             bool
	Dir                  string
	Taskfile             string
	Output               string
	OutputGroupBegin     string
	OutputGroupEnd       string
	OutputGroupErrorOnly bool
	Concurrency          int
	Interval             time.Duration
	Global               bool
	Force                bool
	CliArgs              []string

	fs *pflag.FlagSet
}

func GetArgs(args []string) *Args {
	result := Args{}

	fs := pflag.NewFlagSet(args[0], pflag.ExitOnError)

	fs.BoolVar(&result.Version, "version", false, "Show task-tui version.")
	fs.BoolVarP(&result.Help, "help", "h", false, "Show task-tui usage.")

	fs.BoolVar(&result.EnableSecondLine, "enable-second-line", false, "Show the description next to the task name.")
	fs.BoolVarP(&result.Repeat, "repeat", "r", false, "Repeat execution of the last executed command.")

	fs.StringVar(&result.TaskSort, "sort", "", "(Task) Changes the order of the tasks when listed. [default|alphanumeric|none].")
	fs.BoolVar(&result.Status, "status", false, "(Task) Exits with non-zero exit code if any of the given tasks is not up-to-date.")
	fs.BoolVar(&result.Insecure, "insecure", false, "(Task) Forces Task to download Taskfiles over insecure connections.")
	fs.BoolVarP(&result.Watch, "watch", "w", false, "(Task) Enables watch of the given task.")
	fs.BoolVarP(&result.Verbose, "verbose", "v", false, "(Task) Enables verbose mode.")
	fs.BoolVarP(&result.Silent, "silent", "s", false, "(Task) Disables echoing.")
	fs.BoolVarP(&result.AssumeYes, "yes", "y", false, "(Task) Assume \"yes\" as answer to all prompts.")
	fs.BoolVarP(&result.Parallel, "parallel", "p", false, "(Task) Executes tasks provided on command line in parallel.")
	fs.BoolVarP(&result.Dry, "dry", "n", false, "(Task) Compiles and prints tasks in the order that they would be run, without executing them.")
	fs.BoolVarP(&result.ExitCode, "exit-code", "x", false, "(Task) Pass-through the exit code of the task command.")
	fs.StringVarP(&result.Dir, "dir", "d", "", "(Task) Sets the directory in which Task will execute and look for a Taskfile.")
	fs.StringVarP(&result.Taskfile, "taskfile", "t", "", `(Task) Choose which Taskfile to run. Defaults to "Taskfile.yml".`)
	fs.StringVarP(&result.Output, "output", "o", "", "(Task) Sets output style: [interleaved|group|prefixed].")
	fs.StringVar(&result.OutputGroupBegin, "output-group-begin", "", "(Task) Message template to print before a task's grouped output.")
	fs.StringVar(&result.OutputGroupEnd, "output-group-end", "", "(Task) Message template to print after a task's grouped output.")
	fs.BoolVar(&result.OutputGroupErrorOnly, "output-group-error-only", false, "(Task) Swallow output from successful tasks.")
	fs.IntVarP(&result.Concurrency, "concurrency", "C", 0, "(Task) Limit number of tasks to run concurrently.")
	fs.DurationVarP(&result.Interval, "interval", "I", 0, "(Task) Interval to watch for changes.")
	fs.BoolVarP(&result.Global, "global", "g", false, "(Task) Runs global Taskfile, from $HOME/{T,t}askfile.{yml,yaml}.")
	fs.BoolVarP(&result.Force, "force", "f", false, "(Task) Forces execution even when the task is up-to-date.")
	fs.Parse(args)

	argsLenAtDash := fs.ArgsLenAtDash()
	if argsLenAtDash >= 0 {
		result.CliArgs = fs.Args()[argsLenAtDash:]
	} else {
		result.CliArgs = make([]string, 0, 0)
	}

	result.fs = fs

	return &result
}

func (a *Args) PrintVersion() {
	version := build.Version
	if len(version) > 0 {
		version = "v" + version
	} else {

		version = "dev"
	}
	version = "task-tui " + version
	fmt.Fprintln(os.Stderr, version)
}

func (a *Args) PrintUsage() {
	const usage = `Usage: task-tui [flags...]

Options:
`

	a.PrintVersion()
	fmt.Fprint(os.Stderr, usage)
	a.fs.PrintDefaults()
}

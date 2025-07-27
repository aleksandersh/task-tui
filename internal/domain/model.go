package domain

type Taskfile struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name        string   `json:"name"`
	Task        string   `json:"task"`
	Description string   `json:"desc"`
	Summary     string   `json:"summary"`
	Aliases     []string `json:"aliases"`
}

type Command struct {
	Name    string
	Args    []string
	CliArgs []string
}

func NewCommand(name string, args []string, cliArgs []string) Command {
	return Command{Name: name, Args: args, CliArgs: cliArgs}
}

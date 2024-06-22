package domain

type Taskfile struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name        string   `json:"name"`
	Description string   `json:"desc"`
	Summary     string   `json:"summary"`
	Aliases     []string `json:"aliases"`
}

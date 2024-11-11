package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aleksandersh/task-tui/domain"
)

const fileName = "task_tui_latest_command.json"

type commandDto struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func LoadLatestCommand() (domain.Command, error) {
	content, err := os.ReadFile(latestCommandFileName())
	if err != nil {
		if os.IsNotExist(err) {
			return domain.Command{}, fmt.Errorf("latest command not found")
		}
		return domain.Command{}, fmt.Errorf("failed to load latest command: %w", err)
	}

	var cmd commandDto
	if err = json.Unmarshal(content, &cmd); err != nil {
		return domain.Command{}, fmt.Errorf("failed to deserialize latest command: %w", err)
	}

	return domain.NewCommand(cmd.Name, cmd.Args), nil
}

func SaveLatestCommand(command domain.Command) error {
	content, err := json.Marshal(commandDto{Name: command.Name, Args: command.Args})
	if err != nil {
		return fmt.Errorf("failed to serialize command: %w", err)
	}
	if err = os.WriteFile(latestCommandFileName(), content, 0644); err != nil {
		return fmt.Errorf("failed to write temporary file: %w", err)
	}
	return nil
}

func latestCommandFileName() string {
	return os.TempDir() + "/" + fileName
}

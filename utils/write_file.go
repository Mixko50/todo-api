package utils

import (
	"encoding/json"
	"github.com/Mixko50/todo-api/types"
	"os"
)

func WriteFile(todos []types.Todo) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Mixko50/todo-api/types"
	"io"
	"os"
)

func ReadFile() []types.Todo {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return []types.Todo{}
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	todos := make([]types.Todo, 0)

	err = json.Unmarshal(byteValue, &todos)
	if err != nil {
		return []types.Todo{}
	}

	return todos
}

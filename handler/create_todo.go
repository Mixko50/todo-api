package handler

import (
	"github.com/Mixko50/todo-api/types"
	"github.com/Mixko50/todo-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var todo *types.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	todos := utils.ReadFile()

	if len(todos) == 0 {
		todo.ID = 1
	} else {
		todo.ID = todos[len(todos)-1].ID + 1
	}
	todos = append(todos, *todo)
	if err := utils.WriteFile(todos); err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.HTML(http.StatusCreated, "todos.html", todos)

}

package handler

import (
	"github.com/Mixko50/todo-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteTodo(c *gin.Context) {
	dat := c.Param("ID")

	todos := utils.ReadFile()
	convertNumber, _ := strconv.Atoi(dat)

	for i, v := range todos {
		if v.ID == convertNumber {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	if err := utils.WriteFile(todos); err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.HTML(http.StatusOK, "todos.html", todos)
}

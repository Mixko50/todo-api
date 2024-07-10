package handler

import (
	"github.com/Mixko50/todo-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodos(c *gin.Context) {
	todos := utils.ReadFile()
	c.HTML(http.StatusOK, "todos.html", todos)
}

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Mixko50/todo-api/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	//defer cancel()
	//
	//<-ctx.Done()
	//fmt.Println("Received SIGINT signal. Exiting...")
	//booboo(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//booboo(ctx)
	//defer cancel()
	//<-ctx.Done()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()
	r.LoadHTMLGlob("./*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/todos", handler.GetTodos)
	r.POST("/todos", handler.CreateTodo)
	r.DELETE("/todos/:ID", handler.DeleteTodo)
	r.GET("/ping", pingPongHandler)

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	closedChannel := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("Received signal. Gracefully shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Error:", err)
			}
		}
		close(closedChannel)
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println("Error Serve:", err)
	}

	<-closedChannel
	fmt.Println("Server shutdown successfully")
}

//func booboo(ctx context.Context) {
//	t := time.NewTicker(time.Second * 3)
//	select {
//	case <-ctx.Done():
//		fmt.Println("Timeout")
//	case <-t.C:
//		fmt.Println("Tick")
//	}
//}

func pingPongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

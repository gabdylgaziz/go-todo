package app

import (
	"context"
	"flag"
	"fmt"
	"gotodo/internal/handler"
	"gotodo/internal/repository"
	"gotodo/internal/service/gotodo"
	"gotodo/pkg/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port = "8080"
)

func Run() {
	repositories, err := repository.New(repository.WithMemoryStore())
	if err != nil {
		//TODO: log add
		return
	}

	todoService, err := gotodo.New(gotodo.WithTodoRepository(repositories.Todo))
	if err != nil {
		//TODO: log add
		return
	}

	handlers, err := handler.New(
		handler.Dependencies{
			TodoService: todoService,
		}, handler.WithHTTPHandler(),
	)
	if err != nil {
		//TODO: log add
		return
	}

	servers, err := server.New(
		server.WithHTTPServer(handlers.HTTP, port))
	if err != nil {
		//TODO: log add
		return
	}

	if err = servers.Run(); err != nil {
		//TODO: log add
		return
	}
	fmt.Println("http server started on http://localhost:" + port)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the httpServer gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err = servers.Stop(ctx); err != nil {
		panic(err) // failure/timeout shutting down the httpServer gracefully
	}

	fmt.Println("running cleanup tasks...")

	fmt.Println("server was successful shutdown.")
}

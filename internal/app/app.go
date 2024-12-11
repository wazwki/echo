package app

import (
	"context"
	"example/internal/config"
	"example/internal/http/v1"
	"example/internal/repository"
	"example/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type App struct {
	server *http.Server
}

func New(cfg *config.Config) (*App, error) {

	// logger

	// db

	repository := repository.NewRepository("inited db")
	service := service.NewService(repository)
	controllers := v1.NewControllers(service)

	srv := echo.New()
	api := srv.Group("/api")
	v1 := api.Group("/v1")

	// mux

	srv.Server = &http.Server{
		Addr:              fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		ReadHeaderTimeout: 800 * time.Millisecond,
		ReadTimeout:       800 * time.Millisecond,
	}

	return &App{server: srv}, nil
}

func (a *App) Run() error {

	// migrate

	go func() {
		a.server.Logger.Fatal(a.server.StartServer(a.server.Server))
	}()

	return nil
}

func (a *App) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Server gracefully stopped")

	//db

	return nil
}

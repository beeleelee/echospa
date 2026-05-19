package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Setup
	e := echo.New()
	// Set Bundle MiddleWare
	e.Use(middleware.Gzip())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "/web",       // This is the path to your SPA build folder, the folder that is created from running "npm build"
		Index:  "index.html", // This is the default html page for your SPA
		Browse: false,
		HTML5:  true,
	}))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         ":8080",
		GracefulTimeout: 5 * time.Second,
	}
	if err := sc.Start(ctx, e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

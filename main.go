package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	port := flag.String("port", "", "listen port (default 8080)")
	staticRoot := flag.String("static-root", "", "path to static files root (default ./web)")
	flag.Parse()

	addr := os.Getenv("PORT")
	if *port != "" {
		addr = *port
	}
	if addr == "" {
		addr = ":8080"
	}
	if addr[0] != ':' {
		addr = ":" + addr
	}

	root := os.Getenv("STATIC_ROOT")
	if *staticRoot != "" {
		root = *staticRoot
	}
	if root == "" {
		root = "./web"
	}

	// Setup
	e := echo.New()
	e.Use(middleware.Gzip())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   root,
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         addr,
		GracefulTimeout: 5 * time.Second,
	}
	if err := sc.Start(ctx, e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gouravsingh19/CURD-API/internal/config"
	"github.com/gouravsingh19/CURD-API/internal/http/handlers/student"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database connection
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /api/students", student.New())

	//setup server
	srv := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}
	//start server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	slog.Info("server stopped")
	ctx, canel := context.WithTimeout(context.Background(), 5*time.Second)
	defer canel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", slog.String("error", err.Error()))
	}
	srv.Shutdown(ctx)
}

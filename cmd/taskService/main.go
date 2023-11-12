package main

import (
	"TaskService/internal/app/api/config"
	"TaskService/internal/app/api/handlers"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	timeOut := time.Duration(cfg.TimeOut)
	r, err := handlers.InitRouter(cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HttpServer.Address, cfg.HttpServer.Host),
		Handler: r,
	}

	// create goroutine

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("listen %s\n", err)
		}
	}()

	// block anonymous goroutine
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	select {
	case <-ctx.Done():
		log.Printf("Timeout of %s seconds", strconv.Itoa(cfg.TimeOut))
	}
	log.Println("Server exiting")

}

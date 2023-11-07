package main

import (
	"TaskService/internal/app/api/config"
	"TaskService/internal/app/api/handlers"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()

	r, err := handlers.InitRouter(cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HttpServer.Address, cfg.HttpServer.Host),
		Handler: r,
	}

	// create gorutine

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("listen %s\n", err)
		}
	}()

	//todo do it
	quit := make(chan os.Signal)
	<-quit

	adr := fmt.Sprintf("%s:%s", cfg.HttpServer.Address, cfg.HttpServer.Host)
	err = r.Run(adr)
	if err != nil {
		log.Fatal(err)
	}

}

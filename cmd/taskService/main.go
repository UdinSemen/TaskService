package main

import (
	"TaskService/internal/config"
	"TaskService/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.MustLoad()
	r := gin.Default()

	rout, err := router.InitRouter(cfg, r)
	if err != nil {
		log.Fatal(err)
	}

	rout.AddTask()
	rout.EditTask()
	rout.DeleteTask()
	rout.GetTask()
	rout.AddManyTasks()
	rout.EditManyTasks()

	adr := fmt.Sprintf("%s:%s", cfg.HttpServer.Address, cfg.HttpServer.Host)
	err = r.Run(adr)
	if err != nil {
		log.Fatal(err)
	}

	r.Handler()
}

package handlers

import (
	cfg "TaskService/internal/app/api/config"
	"TaskService/internal/app/api/model"
	"TaskService/internal/app/api/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Rep todo Isn't it good?
var Rep *repository.Repository

type reqForMany []model.Task

func InitRouter(cfg *cfg.Config) (*gin.Engine, error) {
	var err error
	const op = "handlers.InitRouter"
	Rep, err = repository.InitRepository(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	r := gin.New()
	r.Use(gin.Logger())

	r.PUT("/task", AddTask)
	r.PUT("/many_tasks", AddManyTasks)
	r.GET("/task", GetTask)
	r.GET("/many_tasks", GetAllTasks)
	r.POST("/task", EditTask)
	r.POST("/many_tasks", EditManyTask)
	r.DELETE("/task", DeleteTask)
	r.DELETE("/many_tasks", DeleteManyTasks)
	return r, nil
}

func ShowBadReq(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func ReplaceData(new *model.Task, old model.Task) {
	if old.Uri != "" {
		new.Uri = old.Uri
	}
	if old.Type != "" {
		new.Type = old.Type
	}
	if old.CollectionFrequency > 0 {
		new.CollectionFrequency = old.CollectionFrequency
	}
	if old.Description != "" {
		new.Description = old.Description
	}
	if old.Ti != nil {
		new.Ti = old.Ti
	}
}

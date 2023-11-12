package handlers

import (
	cfg "TaskService/internal/app/api/config"
	stor "TaskService/internal/app/api/mesStorage/pgMemStorage"
	"TaskService/internal/app/api/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Rep todo Isn't it good?
var PgStorage *stor.PgMemStorage

type reqForMany []model.Task

func InitRouter(cfg *cfg.Config) (*gin.Engine, error) {
	var err error
	const op = "handlers.InitRouter"
	PgStorage, err = stor.NewPgMemStorage(&cfg.DbConfig)
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

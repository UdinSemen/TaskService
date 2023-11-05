package router

import (
	"TaskService/internal/config"
	"TaskService/internal/model"
	"TaskService/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	GinRouter  *gin.Engine
	Repository *repository.Repository
}

func InitRouter(cfg *config.Config, engine *gin.Engine) (*Router, error) {
	const op = "router.InitRouter"

	rep, err := repository.InitRepository(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Router{
		GinRouter:  engine,
		Repository: rep,
	}, nil
}

func (r *Router) AddTask() {
	r.GinRouter.PUT("/task", func(ctx *gin.Context) {
		var taskReq model.Task

		if err := ctx.BindJSON(&taskReq); err != nil {
			ShowBadReq(ctx, err)
			return
		}

		if err := r.Repository.AddTask(taskReq); err != nil {
			ShowBadReq(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, taskReq)
	})
}

func (r *Router) EditTask() {
	r.GinRouter.POST("/task", func(ctx *gin.Context) {
		var taskReq model.Task

		if err := ctx.BindJSON(&taskReq); err != nil {
			ShowBadReq(ctx, err)
			return
		}

		task, err := r.Repository.GetTask(taskReq.Id)
		if err != nil {
			ShowBadReq(ctx, err)
			return
		}

		// Todo <In my opinion this solution is awful>
		ReplaceData(&task, taskReq)

		ctx.JSON(http.StatusCreated, task)
	})
}

func (r *Router) DeleteTask() {
	r.GinRouter.DELETE("/task", func(ctx *gin.Context) {
		id := ctx.Query("id")
		if err := r.Repository.DeleteTask(id); err != nil {
			ShowBadReq(ctx, err)
			return
		}

		ctx.Status(http.StatusOK)
	})
}

func (r *Router) GetTask() {
	r.GinRouter.GET("/task", func(ctx *gin.Context) {
		id := ctx.Query("id")
		task, err := r.Repository.GetTask(id)
		if err != nil {
			ShowBadReq(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, task)
	})

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

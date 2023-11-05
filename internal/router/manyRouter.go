package router

import (
	"TaskService/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reqForMany []model.Task

func (r *Router) AddManyTasks() {
	r.GinRouter.POST("/many_tasks", func(ctx *gin.Context) {
		var tasks reqForMany

		err := ctx.BindJSON(&tasks)
		if err != nil {
			ShowBadReq(ctx, err)
			return
		}

		err = r.Repository.AddManyTasks(tasks)
		if err != nil {
			ShowBadReq(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, tasks)
	})
}

func (r *Router) EditManyTasks() {
	r.GinRouter.PUT("/many_tasks", func(ctx *gin.Context) {
		var tasks reqForMany
		var addTasks []model.Task

		err := ctx.BindJSON(&tasks)
		if err != nil {
			ShowBadReq(ctx, err)
		}

		for _, task := range tasks {
			t, err := r.Repository.GetTask(task.Id)
			if err != nil {
				ShowBadReq(ctx, err)
				return
			}
			ReplaceData(&t, task)

			addTasks = append(addTasks, t)
		}

		err = r.Repository.AddManyTasks(addTasks)
		if err != nil {
			ShowBadReq(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, addTasks)
	})
}

func (r *Router) DeleteManyTasks() {
	r.GinRouter.DELETE("/many_tasks", func(ctx *gin.Context) {
		var tasks reqForMany

		if err := ctx.BindJSON(&tasks); err != nil {
			ShowBadReq(ctx, err)
			return
		}

		for _, task := range tasks {
			if err := r.Repository.DeleteTask(task.Id); err != nil {
				ShowBadReq(ctx, err)
				return
			}
		}

		ctx.Status(http.StatusOK)
	})
}

func (r *Router) GetAllTasks() {
	r.GinRouter.GET("/many_tasks", func(ctx *gin.Context) {
		tasks, err := r.Repository.GetAllTasks()
		if err != nil {
			ShowBadReq(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, tasks)
	})
}

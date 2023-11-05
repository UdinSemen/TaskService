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

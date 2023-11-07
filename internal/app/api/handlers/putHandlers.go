package handlers

import (
	"TaskService/internal/app/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTask(ctx *gin.Context) {
	var taskReq model.Task

	if err := ctx.BindJSON(&taskReq); err != nil {
		ShowBadReq(ctx, err)
		return
	}

	if err := Rep.AddTask(taskReq); err != nil {
		ShowBadReq(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, taskReq)
}

func AddManyTasks(ctx *gin.Context) {
	var tasks reqForMany

	err := ctx.BindJSON(&tasks)
	if err != nil {
		ShowBadReq(ctx, err)
		return
	}

	err = Rep.AddManyTasks(tasks)
	if err != nil {
		ShowBadReq(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, tasks)
}

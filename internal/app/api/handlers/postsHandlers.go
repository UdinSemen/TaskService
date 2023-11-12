package handlers

import (
	"TaskService/internal/app/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditTask(ctx *gin.Context) {
	var taskReq model.Task

	if err := ctx.BindJSON(&taskReq); err != nil {
		ShowBadReq(ctx, err)
		return
	}

	task, err := PgStorage.GetTask(ctx, taskReq.Id)
	if err != nil {
		ShowBadReq(ctx, err)
		return
	}

	ReplaceData(&task, taskReq)

	ctx.JSON(http.StatusCreated, task)
}

func EditManyTask(ctx *gin.Context) {
	var tasks reqForMany
	var addTasks []model.Task

	err := ctx.BindJSON(&tasks)
	if err != nil {
		ShowBadReq(ctx, err)
	}

	for _, task := range tasks {
		t, err := PgStorage.GetTask(ctx, task.Id)
		if err != nil {
			ShowBadReq(ctx, err)
			return
		}
		ReplaceData(&t, task)

		addTasks = append(addTasks, t)
	}

	err = PgStorage.AddManyTasks(ctx, addTasks)
	if err != nil {
		ShowBadReq(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, addTasks)
}

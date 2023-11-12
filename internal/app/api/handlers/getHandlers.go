package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTask(ctx *gin.Context) {
	id := ctx.Query("id")
	task, err := PgStorage.GetTask(ctx, id)
	if err != nil {
		ShowBadReq(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func GetAllTasks(ctx *gin.Context) {
	tasks, err := PgStorage.GetAllTasks(ctx)
	if err != nil {
		ShowBadReq(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

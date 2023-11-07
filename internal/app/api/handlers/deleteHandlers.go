package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteTask(ctx *gin.Context) {
	id := ctx.Query("id")
	if err := Rep.DeleteTask(id); err != nil {
		ShowBadReq(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func DeleteManyTasks(ctx *gin.Context) {
	var tasks reqForMany

	if err := ctx.BindJSON(&tasks); err != nil {
		ShowBadReq(ctx, err)
		return
	}

	for _, task := range tasks {
		if err := Rep.DeleteTask(task.Id); err != nil {
			ShowBadReq(ctx, err)
			return
		}
	}

	ctx.Status(http.StatusOK)
}

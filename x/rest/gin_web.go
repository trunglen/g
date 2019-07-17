package rest

import (
	"github.com/gin-gonic/gin"
)

const STATUS_OK = 200

type JsonRender struct {
}

func (r *JsonRender) GetUserID(ctx *gin.Context) string {
	userID, exist := ctx.Get("user_id")
	if exist {
		return userID.(string)
	}
	return ctx.Query("user_id")
}

func (r *JsonRender) SendData(ctx *gin.Context, data interface{}) {
	ctx.JSON(STATUS_OK, map[string]interface{}{
		"data":   data,
		"status": "success",
	})
}

func (r *JsonRender) SendDataCount(ctx *gin.Context, data interface{}, count int) {
	ctx.JSON(STATUS_OK, map[string]interface{}{
		"data":   data,
		"count":  count,
		"status": "success",
	})
}

func (r *JsonRender) SendArrayData(ctx *gin.Context, data interface{}) {
	if data == nil {
		ctx.JSON(STATUS_OK, map[string]interface{}{
			"data":   make([]string, 0),
			"status": "success",
		})
	} else {
		ctx.JSON(STATUS_OK, map[string]interface{}{
			"data":   data,
			"status": "success",
		})
	}
}

func (r *JsonRender) Success(ctx *gin.Context) {
	ctx.JSON(STATUS_OK, map[string]interface{}{
		"data":   nil,
		"status": "success",
	})
}

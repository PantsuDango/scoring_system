package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scoring_system/model/tables"
)

func (Controller Controller) ListUser(ctx *gin.Context, user tables.User) {
	// 如果不是主账号
	if user.Type != 1 {
		JSONFail(ctx, http.StatusOK, AccessDeny, "user type error.", gin.H{
			"Code":    "InvalidJSON",
			"Message": "user type error",
		})
		return
	}

	user_info := Controller.ScoringDB.QueryAllUser()

	JSONSuccess(ctx, http.StatusOK, user_info)
}

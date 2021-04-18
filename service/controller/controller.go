package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scoring_system/model/tables"
)

func (Controller Controller) InitTest(ctx *gin.Context, user tables.User) {
	JSONSuccess(ctx, http.StatusOK, user)
}

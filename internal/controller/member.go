package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
}

func NewMemberController() *MemberController {
	return &MemberController{}
}

func (controller MemberController) CreateMockDataForMembers(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateMockDataForMembers",
	})
}

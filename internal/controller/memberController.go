package controller

import (
	"contract/internal/service/memberService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
	memberServ *memberService.MemberService
}

func NewMemberController(memberService *memberService.MemberService) *MemberController {
	return &MemberController{
		memberServ: memberService,
	}
}

func (m MemberController) CreateMockDataForMembers(ctx *gin.Context) {

	// 產生 1000 筆 member 資料
	m.memberServ.CreateMockDataForMembers()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

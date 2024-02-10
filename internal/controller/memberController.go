package controller

import (
	"contract/internal/data/dto"
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

func (m MemberController) UpdateMember(ctx *gin.Context) {

	var memberID = ctx.Param("id")
	var requestDTO dto.UpdateMemberRequestDTO

	err := ctx.BindJSON(&requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
	}

	// 更新 member 資料
	m.memberServ.UpdateMember(memberID, requestDTO)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func (m MemberController) CreateMember(ctx *gin.Context) {

	var requestDTO dto.PostMemberRequestDTO

	err := ctx.BindJSON(&requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
	}

	// 新增 member 資料
	m.memberServ.CreateMember(requestDTO)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func (m MemberController) GetMembersTransactionsYearly(ctx *gin.Context) {

	// 取得所有 member 的年度交易資料
	result := m.memberServ.GetMembersTransactionsYearly()

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (m MemberController) GetMembers(context *gin.Context) {

	// 取得所有 member 資料
	result := m.memberServ.GetMembers()

	context.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

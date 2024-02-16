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

func (m MemberController) GetMembers(ctx *gin.Context) {

	// 取得所有 member 資料
	result := m.memberServ.GetMembers()

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (m MemberController) GetMemberTransactions(ctx *gin.Context) {

	id := ctx.Param("id")
	startDateTime := ctx.Query("start_datetime")
	endDateTime := ctx.Query("end_datetime")

	result := m.memberServ.GetMembersTransactions(id, startDateTime, endDateTime)

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (m MemberController) TestCICD(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You are testing CICD - {1.1}",
	})
}

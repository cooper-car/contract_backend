package http

import (
	"contract/internal/controller"
	"contract/internal/repository"
	"contract/internal/repository/borrowFeeRepository"
	"contract/internal/repository/memberRepository"
	"contract/internal/service/memberService"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db := repository.GetDB()

	memberRepo := memberRepository.NewMemberRepository(db)
	borrowFeeRepo := borrowFeeRepository.NewBorrowFeeRepository(db)
	memberServ := memberService.NewMemberService(memberRepo, borrowFeeRepo)
	memberController := controller.NewMemberController(memberServ)

	router.POST("/api/member/mock/data", memberController.CreateMockDataForMembers)

	return router
}

package http

import (
	"contract/internal/controller"
	"contract/internal/repository"
	"contract/internal/repository/borrowFeeRepository"
	"contract/internal/repository/memberRepository"
	"contract/internal/service/memberService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	db := repository.GetDB()

	memberRepo := memberRepository.NewMemberRepository(db)
	borrowFeeRepo := borrowFeeRepository.NewBorrowFeeRepository(db)
	memberServ := memberService.NewMemberService(memberRepo, borrowFeeRepo)
	memberController := controller.NewMemberController(memberServ)

	router.POST("/api/member/mock/data", memberController.CreateMockDataForMembers)
	router.GET("/api/members", memberController.GetMembers)
	router.PATCH("/api/member/:id", memberController.UpdateMember)
	router.POST("/api/member", memberController.CreateMember)
	router.GET("/api/members/:id/transactions", memberController.GetMemberTransactions)
	router.GET("/api/test-cicd", memberController.TestCICD)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 接收客戶端發送的origin
		c.Header("Access-Control-Allow-Origin", "*")

		// 允許客戶端傳遞校驗信息比如 cookie
		c.Header("Access-Control-Allow-Credentials", "true")

		// 允許跨域設置可以返回其他子段，可以自定義字段
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Account, URL")

		// 服務器支持的所有跨域請求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")

		// 設置緩存時間
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

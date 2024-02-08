package http

import (
	"contract/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	memberController := controller.NewMemberController()
	router.POST("/api/member/mock/data", memberController.CreateMockDataForMembers)

	return router
}

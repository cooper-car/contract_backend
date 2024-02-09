package main

import (
	"contract/internal/repository"
	"contract/pkg/http"
)

func main() {
	// 初始化資料庫(Mysql)
	repository.Setup()

	router := http.SetupRouter()
	router.Run(":8080")
}

package main

import "contract/pkg/http"

func main() {
	router := http.SetupRouter()
	router.Run(":8080")
}

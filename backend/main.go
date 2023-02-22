package main

import (
	"Karaokifyy/backend/search"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	search.Init(router)
	router.Run("localhost:8080")
}

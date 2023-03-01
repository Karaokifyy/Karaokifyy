package main

import (
	"Karaokifyy/backend/search"

	"github.com/gin-gonic/gin"
)

func main() {
	// search.SearchBySong("Hyperballad")

	router := gin.Default()
	search.Init(router)
	router.Run("localhost:40")
}

package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/Karaokifyy/Karaokifyy/app/karaokifyy_api"
	"github.com/Karaokifyy/Karaokifyy/app/spotify_api"
	"github.com/Karaokifyy/Karaokifyy/BackEnd/src/api/routes/router"
)

func main() {
	// TODO: Router runs configs
	router.Run()
}



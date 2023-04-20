package view

import (
	"encoding/json"
	"net/http"
)

func EncodeSongList(w http.ResponseWriter, s interface{}) {
	json.NewEncoder(w).Encode(s)
}

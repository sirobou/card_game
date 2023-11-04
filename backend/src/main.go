package main

import (
	"casino/handlers/actionHandler"
	"casino/handlers/createNewPlayerHandler"
	"casino/handlers/getPlayerHandler"
	"casino/model/game"
	"casino/model/player"
	"net/http"
)

func main() {
	game := game.CreateNewGame()

	http.HandleFunc("/api/game/join", func(w http.ResponseWriter, r *http.Request) {
		createNewPlayerHandler.CreateNewPlayerHandler(w, r, game)
	})
	http.HandleFunc("/api/player/action", func(w http.ResponseWriter, r *http.Request) {
		actionHandler.ActionHandler(w, r, game)
	})
	http.HandleFunc("/api/game/player", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		player_id := queryParams.Get("id")
		getPlayerHandler.GetPlayerHandler(w, r, player.PlayerId(player_id), game)
	})
	http.ListenAndServe("localhost:8080", nil)
}

package main

import (
	"casino/handlers/actionHandler"
	"casino/handlers/createNewPlayerHandler"
	getgamehandler "casino/handlers/getGameHandler.go"
	"casino/handlers/getPlayerHandler"
	"casino/model/game"
	"casino/model/player"
	"fmt"
	"net/http"
)

func main() {
	game := game.CreateNewGame()
	game.Deck.Shuffle()

	http.HandleFunc("/api/game/join", func(w http.ResponseWriter, r *http.Request) {
		createNewPlayerHandler.CreateNewPlayerHandler(w, r, game)
		for _, p := range game.Players {
			fmt.Println(p.Name)
		}
	})
	http.HandleFunc("/api/player/action", func(w http.ResponseWriter, r *http.Request) {
		actionHandler.ActionHandler(w, r, game)
	})
	http.HandleFunc("/api/game/player", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		playerId := queryParams.Get("id")
		getPlayerHandler.GetPlayerHandler(w, r, player.PlayerId(playerId), game)
	})
	http.HandleFunc("/api/game", func(w http.ResponseWriter, r *http.Request) {
		getgamehandler.GetGameHandler(w, r, game)
	})
	http.ListenAndServe("localhost:8080", nil)
}

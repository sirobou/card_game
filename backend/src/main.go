package main

import (
	"casino/handlers/actionHandler"
	"casino/handlers/createNewPlayerHandler"
	"casino/handlers/getPlayerHandler"
	"casino/handlers/getRoundHandler.go"
	"casino/handlers/startRoundHandler"
	"casino/model/lobby"
	"casino/model/player"
	"fmt"
	"net/http"
)

func main() {
	lobby := lobby.NewLobby()
	http.HandleFunc("/api/lobby/join", func(w http.ResponseWriter, r *http.Request) {
		createNewPlayerHandler.CreateNewPlayerHandler(w, r, lobby)
		fmt.Println(lobby.Players)
	})
	http.HandleFunc("/api/lobby/start", func(w http.ResponseWriter, r *http.Request) {
		startRoundHandler.StartRoundHandler(w, r, lobby)
		fmt.Println(lobby.Round.Players)
	})
	http.HandleFunc("/api/round/action", func(w http.ResponseWriter, r *http.Request) {
		actionHandler.ActionHandler(w, r, lobby.Round)
	})
	http.HandleFunc("/api/round/player", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		playerId := queryParams.Get("id")
		getPlayerHandler.GetPlayerHandler(w, r, player.PlayerId(playerId), lobby.Round)
	})
	http.HandleFunc("/api/round", func(w http.ResponseWriter, r *http.Request) {
		getRoundHandler.GetRoundHandler(w, r, lobby.Round)
	})
	http.ListenAndServe("localhost:8080", nil)
}

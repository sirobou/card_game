package main

import (
	"casino/handlers/actionHandler"
	"casino/handlers/createNewPlayerHandler"
	"casino/handlers/getLobbyHandler"
	"casino/handlers/getPlayerHandler"
	"casino/handlers/getResultHandler"
	"casino/handlers/getRoundHandler"
	"casino/handlers/resetLobbyHandler"
	"casino/handlers/startRoundHandler"
	"casino/model/lobby"
	"casino/model/player"
	"net/http"
)

func main() {
	lobby := lobby.NewLobby()
	http.HandleFunc("/api/lobby/join", func(w http.ResponseWriter, r *http.Request) {
		createNewPlayerHandler.CreateNewPlayerHandler(w, r, lobby)
	})
	http.HandleFunc("/api/lobby/start", func(w http.ResponseWriter, r *http.Request) {
		startRoundHandler.StartRoundHandler(w, r, lobby)
	})
	http.HandleFunc("/api/lobby", func(w http.ResponseWriter, r *http.Request) {
		getLobbyHandler.GetLobbyHandler(w, r, lobby)
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
	http.HandleFunc("/api/round/result", func(w http.ResponseWriter, r *http.Request) {
		getResultHandler.GetResultHandler(w, r, lobby.Round)
	})
	http.HandleFunc("/api/lobby/reset", func(w http.ResponseWriter, r *http.Request) {
		resetLobbyHandler.ResetLobbyHandler(w, r, lobby)
	})
	http.ListenAndServe("0.0.0.0:80", nil)
}

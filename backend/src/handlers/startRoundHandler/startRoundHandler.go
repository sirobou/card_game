package startRoundHandler

import (
	"casino/model/lobby"
	"net/http"
)

func StartRoundHandler(w http.ResponseWriter, r *http.Request, lobby *lobby.Lobby) {
	if r.Method == http.MethodGet {
		lobby.Round = lobby.StartRound()
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
}

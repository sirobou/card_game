package resetLobbyHandler

import (
	"casino/model/lobby"
	"net/http"
)

func ResetLobbyHandler(w http.ResponseWriter, r *http.Request, l *lobby.Lobby) {
	if r.Method == http.MethodGet {
		*l = *lobby.NewLobby()
	}
}

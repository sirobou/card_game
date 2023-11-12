package getLobbyHandler

import (
	"casino/handlers/serializer"
	"casino/model/lobby"
	"fmt"
	"net/http"
)

func GetLobbyHandler(w http.ResponseWriter, r *http.Request, l *lobby.Lobby) {
	if r.Method == http.MethodGet {
		CreateJsonResponseFromLobby(w, *l)
	}
}

func CreateJsonResponseFromLobby(w http.ResponseWriter, l lobby.Lobby) {
	var lobbySerializer serializer.Serializer[lobby.Lobby] = serializer.NewLobbySerialiezr()
	jsonData, err := lobbySerializer.Serialize(l)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

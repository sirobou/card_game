package getLobbyHandler

import (
	"casino/handlers/converter"
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
	playerConverter := converter.NewPlayerConverter()
	var lobbySerializer serializer.Serializer[lobby.Lobby] = serializer.NewLobbySerialiezr(playerConverter)
	jsonData, err := lobbySerializer.Serialize(l)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Origin", "https://sirobou-casino.netlify.app")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

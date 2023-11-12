package createNewPlayerHandler

import (
	"casino/consts"
	"casino/handlers/converter"
	"casino/handlers/serializer"
	"casino/model/lobby"
	"casino/model/player"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateNewPlayerHandler(w http.ResponseWriter, r *http.Request, l *lobby.Lobby) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		if len(l.Players) == consts.MAX_PLAYER {
			http.Error(w, "The number of players is limited to 4", http.StatusBadRequest)
			return
		}

		var requestData RequestData
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
			return
		}
		name := requestData.Name
		newPlayer, err := player.NewPlayer(name, l.GenerateUniquePlayerId())
		if err != nil {
			fmt.Println("Error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		l.Join(newPlayer)
		CreateJsonResponseFromNewPlayer(w, *newPlayer)
	}
}

func CreateJsonResponseFromNewPlayer(w http.ResponseWriter, Player player.Player) {

	playerConverter := converter.NewPlayerConverter()
	var playerSerializer serializer.Serializer[player.Player] = serializer.NewPlayerSerializer(playerConverter)
	jsonData, err := playerSerializer.Serialize(Player)

	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

type RequestData struct {
	Name string `json:"name"`
}

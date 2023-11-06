package createNewPlayerHandler

import (
	"casino/handlers/converter"
	"casino/handlers/serializer"
	"casino/model/game"
	"casino/model/player"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateNewPlayerHandler(w http.ResponseWriter, r *http.Request, g *game.Game) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var requestData RequestData
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
			return
		}
		name := requestData.Name
		newPlayer, err := player.NewPlayer(g.Deck, name, g.GenerateUniquePlayerId())
		if err != nil {
			fmt.Println("Error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		g.Join(newPlayer)
		fmt.Println(*newPlayer.Hand[1])
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
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

type RequestData struct {
	Name string `json:"name"`
}
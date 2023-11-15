package getPlayerHandler

import (
	"casino/handlers/converter"
	"casino/handlers/serializer"
	"casino/model/player"
	"casino/model/round"
	"fmt"
	"net/http"
)

func GetPlayerHandler(w http.ResponseWriter, r *http.Request, playerId player.PlayerId, round *round.Round) {
	if r.Method == http.MethodGet {
		player, err := round.GetPlayer(playerId)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			CreateJsonResponseFromPlayer(w, *player)
		}
	}
}

func CreateJsonResponseFromPlayer(w http.ResponseWriter, Player player.Player) {
	playerConverter := converter.NewPlayerConverter()
	var playerSerializer serializer.Serializer[player.Player] = serializer.NewPlayerSerializer(playerConverter)
	jsonData, err := playerSerializer.Serialize(Player)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Origin", "https://sirobou-casino.netlify.app")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

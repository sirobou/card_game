package getPlayerHandler

import (
	"casino/handlers/converter"
	"casino/handlers/serializer"
	"casino/model/game"
	"casino/model/player"
	"fmt"
	"net/http"
)

func GetPlayerHandler(w http.ResponseWriter, r *http.Request, playerId player.PlayerId, Game *game.Game) {
	if r.Method == http.MethodGet {
		fmt.Println(playerId)
		player, err := Game.GetPlayer(playerId)
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
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

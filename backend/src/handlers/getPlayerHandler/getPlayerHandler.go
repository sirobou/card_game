package getPlayerHandler

import (
	"casino/model/game"
	"casino/model/player"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPlayerHandler(w http.ResponseWriter, r *http.Request, playerId player.PlayerId, Game *game.Game) {
	if r.Method == http.MethodGet {
		player, err := Game.GetPlayer(playerId)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			CreateJsonResponseFromPlayer(w, *player)
		}
	}
}

func CreateJsonResponseFromPlayer(w http.ResponseWriter, Player player.Player) {
	jsonData, err := json.Marshal(Player)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

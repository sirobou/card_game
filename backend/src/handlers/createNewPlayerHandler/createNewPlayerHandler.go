package createNewPlayerHandler

import (
	"casino/model/game"
	"casino/model/player"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateNewPlayerHandler(w http.ResponseWriter, r *http.Request, g *game.Game) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		new_player, err := player.NewPlayer(g.Deck, name, g.GenerateUniquePlayerId())
		if err != nil {
			fmt.Println("Error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		CreateJsonResponseFromNewPlayer(w, *new_player)
	}
}

func CreateJsonResponseFromNewPlayer(w http.ResponseWriter, Player player.Player) {
	jsonData, err := json.Marshal(Player)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

package getgamehandler

import (
	"casino/handlers/converter"
	"casino/handlers/serializer"
	"casino/model/game"
	"fmt"
	"net/http"
)

func GetGameHandler(w http.ResponseWriter, r *http.Request, Game *game.Game) {
	if r.Method == http.MethodGet {
		CreateJsonResponseFromGame(w, Game)
	}
}

func CreateJsonResponseFromGame(w http.ResponseWriter, Game *game.Game) {
	playerConverter := converter.NewPlayerConverter()
	gameConverter := converter.NewGameConverter(playerConverter)
	var gameSerializer serializer.Serializer[game.Game] = serializer.NewGameSerializer(gameConverter)
	jsonData, err := gameSerializer.Serialize(*Game)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

package createnewgamehandler

import (
	"casino/model/game"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateNewGameHandler(w http.ResponseWriter, r *http.Request) *game.Game {
	if r.Method == http.MethodGet {
		NewGame := game.CreateNewGame()

		CreateJsonResponseFromNewGame(w, *NewGame)
		return NewGame
	}
	return nil
}

func CreateJsonResponseFromNewGame(w http.ResponseWriter, game game.Game) {
	jsonData, err := json.Marshal(game)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

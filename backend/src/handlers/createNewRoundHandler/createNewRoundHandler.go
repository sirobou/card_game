package createNewRoundHandler

import (
	"casino/model/lobby"
	"casino/model/round"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateNewRoundHandler(w http.ResponseWriter, r *http.Request, l lobby.Lobby) *round.Round {
	if r.Method == http.MethodGet {
		NewRound := round.NewRound(l.Players)

		CreateJsonResponseFromNewGame(w, *NewRound)
		return NewRound
	}
	return nil
}

func CreateJsonResponseFromNewGame(w http.ResponseWriter, round round.Round) {
	jsonData, err := json.Marshal(round)
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

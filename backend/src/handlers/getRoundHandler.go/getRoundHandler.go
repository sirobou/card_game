package getRoundHandler

import (
	"casino/handlers/converter"
	"casino/handlers/serializer"
	"casino/model/round"
	"fmt"
	"net/http"
)

func GetRoundHandler(w http.ResponseWriter, r *http.Request, currentRound *round.Round) {
	if r.Method == http.MethodGet {
		CreateJsonResponseFromGame(w, currentRound)
	}
}

func CreateJsonResponseFromGame(w http.ResponseWriter, currentRound *round.Round) {
	playerConverter := converter.NewPlayerConverter()
	gameConverter := converter.NewRoundConverter(playerConverter)
	var gameSerializer serializer.Serializer[round.Round] = serializer.NewRoundSerializer(gameConverter)
	jsonData, err := gameSerializer.Serialize(*currentRound)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

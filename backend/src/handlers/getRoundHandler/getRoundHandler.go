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
	dealerConverter := converter.NewDealerConverter()
	roundConverter := converter.NewRoundConverter(playerConverter, dealerConverter)
	var gameSerializer serializer.Serializer[round.Round] = serializer.NewRoundSerializer(roundConverter)
	jsonData, err := gameSerializer.Serialize(*currentRound)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Origin", "https://sirobou-casino.netlify.app")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

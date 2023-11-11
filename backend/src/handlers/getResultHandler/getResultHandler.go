package getResultHandler

import (
	"casino/handlers/converter"
	"casino/handlers/serializer"
	"casino/model/round"
	"fmt"
	"net/http"
)

func GetResultHandler(w http.ResponseWriter, r *http.Request, currentRound *round.Round) {
	if r.Method == http.MethodGet {
		if !currentRound.IsEndRound() {
			http.Error(w, "ROUND IS NOT OVER", http.StatusBadRequest)
			return
		}
		CreateJsonResponseFromResult(w, currentRound)
	}
}

func CreateJsonResponseFromResult(w http.ResponseWriter, currentRound *round.Round) {
	playerConverter := converter.NewPlayerConverter()
	resultPlayerConverter := converter.NewResultPlayerConverter(playerConverter, currentRound)
	resultDealerConverter := converter.NewReulstDealerConverter()
	resultRoundConverter := converter.NewResultRoundConverter(resultPlayerConverter, resultDealerConverter)
	var resultRoundSerializer serializer.Serializer[round.Round] = serializer.NewResultRoundSerializer(resultRoundConverter)
	jsonData, err := resultRoundSerializer.Serialize(*currentRound)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
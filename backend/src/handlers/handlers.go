package handlers

import (
	"casino/model/card"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateJsonResponseFromCard(w http.ResponseWriter, Card *card.Card) {
	jsonData, err := json.Marshal(Card)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
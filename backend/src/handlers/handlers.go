package handlers

import (
	"casino/model/card"
	"casino/model/deck"
	"casino/model/player"
	"casino/model/player/id"
	"encoding/json"
	"fmt"
	"net/http"
)

func DrawCardHandler(w http.ResponseWriter, r *http.Request, d *deck.Deck) {
	if r.Method == http.MethodGet {
		new_card := d.Draw()
		CreateJsonResponseFromCard(w, new_card)
	}
}

func CreateNewPlayerHandler(w http.ResponseWriter, r *http.Request, d *deck.Deck) {
	if r.Method == http.MethodGet {
		new_player := player.NewPlayer(d)
		CreateJsonResponseFromPlayer(w, new_player)
	}
}

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

func CreateJsonResponseFromPlayer(w http.ResponseWriter, Player player.Player) {
	var hand []card.ConvertedCard
	for _, c := range Player.Hand {
		hand = append(hand, c.Convert())
	}

	response := struct {
		Id         id.PlayerId
		Hand       []card.ConvertedCard
		Hand_total int
	}{
		Id:         Player.Id,
		Hand:       hand,
		Hand_total: Player.GetHandTotal(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

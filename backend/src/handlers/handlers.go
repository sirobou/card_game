package handlers

import (
	"casino/deck"
	"casino/player"
	"encoding/json"
	"fmt"
	"net/http"
)

func Drawcardhandler(w http.ResponseWriter, r *http.Request, d *deck.Deck) {
	if r.Method == http.MethodGet {
		new_card := d.Draw()
		CreateJsonResponseFromCard(w, new_card)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintln(w, "POSTメソッドが実行されました。")
	} else {
		http.Error(w, "POSTメソッドのリクエストが必要です", http.StatusMethodNotAllowed)
	}
}

func Createnewplayerhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		new_player := player.NewPlayer()
		CreateJsonResponseFromPlayer(w, new_player)
	}
}

func CreateJsonResponseFromCard(w http.ResponseWriter, Card deck.Card) {
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
	jsonData, err := json.Marshal(Player)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

package main

import (
	"casino/handlers"
	"casino/model/deck"
	"net/http"
)

func main() {
	new_Deck := deck.NewDeck()
	new_Deck.Shuffle()

	http.HandleFunc("/create_new_player", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateNewPlayerHandler(w, r, new_Deck)
	})
	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		handlers.DrawCardHandler(w, r, new_Deck)
	})

	http.ListenAndServe("localhost:8080", nil)
}

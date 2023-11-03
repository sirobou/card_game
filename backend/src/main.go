package main

import (
	"casino/deck"
	"casino/handlers"
	"net/http"
)

func main() {
	new_Deck := deck.NewDeck()
	new_Deck.Shuffle()

	http.HandleFunc("/create_new_player", func(w http.ResponseWriter, r *http.Request) {
		handlers.Createnewplayerhandler(w, r)
	})
	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		handlers.Drawcardhandler(w, r, &new_Deck)
	})

	http.ListenAndServe("localhost:8080", nil)
}

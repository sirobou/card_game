package actionHandler

import (
	"casino/model/game"
	"casino/model/player"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ActionHandler(w http.ResponseWriter, r *http.Request, g *game.Game) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var requestData RequestData
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}
	action := requestData.Action
	id := requestData.Id

	switch action {
	case "hit":
		err := g.Hit(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	case "fold":
		err := g.Fold(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, "action only accept 'hit' or 'fold'", http.StatusBadRequest)
		return
	}
}

type RequestData struct {
	Id     player.PlayerId `json:"id"`
	Action string          `json:"action"`
}

package actionHandler

import (
	"casino/model/player"
	"casino/model/round"
	"casino/model/round/action"
	"encoding/json"
	"io"
	"net/http"
)

func ActionHandler(w http.ResponseWriter, r *http.Request, currentRound *round.Round) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var requestData RequestData
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}
	requestAction := requestData.Action
	id := requestData.Id

	err = currentRound.IsPlayerTurn(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch requestAction {
	case "hit":
		err := currentRound.Action(id, action.Hit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	case "stand":
		err := currentRound.Action(id, action.Stand)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, "action only accept 'hit' or 'stand'", http.StatusBadRequest)
		return
	}
}

type RequestData struct {
	Id     player.PlayerId `json:"id"`
	Action string          `json:"action"`
}

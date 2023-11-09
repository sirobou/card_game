package lobby

import (
	"casino/model/player"
	"casino/model/round"
	"strconv"
)

type Lobby struct {
	Round   *round.Round
	Players []*player.Player
}

func NewLobby() *Lobby {
	lobby := Lobby{
		Round:   nil,
		Players: []*player.Player{},
	}
	return &lobby
}

func (l *Lobby) GenerateUniquePlayerId() player.PlayerId {
	for i := 0; ; i++ {
		isUnique := true
		for _, p := range l.Players {
			if p.Id == player.PlayerId(strconv.Itoa(i)) {
				isUnique = false
				break
			}
		}
		if isUnique {
			return player.PlayerId(strconv.Itoa(i))
		}
	}
}

func (l *Lobby) Join(p *player.Player) {
	l.Players = append(l.Players, p)
}

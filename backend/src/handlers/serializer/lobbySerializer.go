package serializer

import (
	"casino/handlers/converter"
	"casino/model/lobby"
	"casino/model/player"
	"encoding/json"
	"fmt"
)

type LobbySerializer struct {
	playerConverter converter.Converter[player.Player, converter.ConvertedPlayer]
}

func (ls *LobbySerializer) Serialize(l lobby.Lobby) ([]byte, error) {
	players := []*converter.ConvertedPlayer{}
	for _, player := range l.Players {
		players = append(players, ls.playerConverter.Convert(player))
	}
	return json.Marshal(players)
}

func (ls *LobbySerializer) Print(l lobby.Lobby) {
	fmt.Println(l.Players)
}

func NewLobbySerialiezr(playerConverter converter.Converter[player.Player, converter.ConvertedPlayer]) *LobbySerializer {
	return &LobbySerializer{playerConverter: playerConverter}
}

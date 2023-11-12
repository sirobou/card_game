package serializer

import (
	"casino/model/lobby"
	"encoding/json"
	"fmt"
)

type LobbySerializer struct {
}

func (ls *LobbySerializer) Serialize(l lobby.Lobby) ([]byte, error) {
	return json.Marshal(l.Players)
}

func (ls *LobbySerializer) Print(l lobby.Lobby) {
	fmt.Println(l.Players)
}

func NewLobbySerialiezr() *LobbySerializer {
	return &LobbySerializer{}
}

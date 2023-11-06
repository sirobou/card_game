package serializer

import (
	"casino/handlers/converter"
	"casino/model/player"
	"encoding/json"
	"fmt"
)

type PlayerSerializer struct {
	playerConverter converter.Converter[player.Player, converter.ConvertedPlayer]
}

func (p *PlayerSerializer) Serialize(pl player.Player) ([]byte, error) {
	return json.Marshal(p.playerConverter.Convert(&pl))
}

func (p *PlayerSerializer) Print(pl player.Player) {
	fmt.Println(p.playerConverter.Convert(&pl))
}

func NewPlayerSerializer(playerConverter converter.Converter[player.Player, converter.ConvertedPlayer]) *PlayerSerializer {
	return &PlayerSerializer{
		playerConverter,
	}
}

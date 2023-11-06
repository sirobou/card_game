package serializer

import (
	"casino/handlers/converter"
	"casino/model/game"
	"encoding/json"
	"fmt"
)

type GameSerializer struct {
	gameConverter converter.Converter[game.Game, converter.ConvertedGame]
}

func (gs *GameSerializer) Serialize(g game.Game) ([]byte, error) {
	return json.Marshal(gs.gameConverter.Convert(&g))
}

func (gs *GameSerializer) Print(g game.Game) {
	fmt.Println(gs.gameConverter.Convert(&g))
}

func NewGameSerializer(gameConverter converter.Converter[game.Game, converter.ConvertedGame]) *GameSerializer {
	return &GameSerializer{
		gameConverter,
	}
}

// func NewConvertedGame(convertedPlayer ConvertedPlayer) *ConvertedGame {
// 	return &ConvertedGame{
// 		Players: convertedPlayer,
// 	}
// }

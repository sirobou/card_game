package serializer

import (
	"casino/handlers/converter"
	"casino/model/round"
	"encoding/json"
	"fmt"
)

type RoundSerializer struct {
	roundConverter converter.Converter[round.Round, converter.ConvertedRound]
}

func (rs *RoundSerializer) Serialize(r round.Round) ([]byte, error) {
	return json.Marshal(rs.roundConverter.Convert(&r))
}

func (rs *RoundSerializer) Print(r round.Round) {
	fmt.Println(rs.roundConverter.Convert(&r))
}

func NewRoundSerializer(roundConverter converter.Converter[round.Round, converter.ConvertedRound]) *RoundSerializer {
	return &RoundSerializer{
		roundConverter,
	}
}

// func NewConvertedGame(convertedPlayer ConvertedPlayer) *ConvertedGame {
// 	return &ConvertedGame{
// 		Players: convertedPlayer,
// 	}
// }

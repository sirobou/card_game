package serializer

import (
	"casino/handlers/converter"
	"casino/model/round"
	"encoding/json"
	"fmt"
)

type ResultRoundSerializer struct {
	ResultRoundConverter converter.Converter[round.Round, converter.ConvertedResultRound]
}

func (rrs *ResultRoundSerializer) Serialize(r round.Round) ([]byte, error) {
	return json.Marshal(rrs.ResultRoundConverter.Convert(&r))
}

func (rrs *ResultRoundSerializer) Print(r round.Round) {
	fmt.Println(rrs.ResultRoundConverter.Convert(&r))
}

func NewResultRoundSerializer(resultRoundConverter converter.Converter[round.Round, converter.ConvertedResultRound]) *ResultRoundSerializer {
	return &ResultRoundSerializer{
		ResultRoundConverter: resultRoundConverter,
	}
}

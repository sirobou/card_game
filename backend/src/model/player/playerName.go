package player

import "errors"

type PlayerName struct {
	Value string
}

func NewPlayerName(value string) (PlayerName, error) {
	if len(value) < 1 || len(value) > 16 {
		return PlayerName{""}, errors.New("invalid name length")
	}
	return PlayerName{value}, nil
}

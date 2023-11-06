package player

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type PlayerName struct {
	Value string
}

func NewPlayerName(value string) (PlayerName, error) {
	if utf8.RuneCountInString(value) < 1 || utf8.RuneCountInString(value) > 16 {
		fmt.Println(len(value))
		return PlayerName{""}, errors.New("invalid name length")
	}
	return PlayerName{value}, nil
}

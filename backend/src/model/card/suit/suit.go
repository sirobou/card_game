package suit

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Spades
	Clubs
)

func ConvertToString(suit Suit) string {
	switch suit {
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Spades:
		return "Spades"
	default:
		return "Clubs"
	}
}

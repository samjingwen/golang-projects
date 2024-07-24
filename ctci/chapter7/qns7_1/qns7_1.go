package qns7_1

type Suit int
type Value int

const (
	Diamonds Suit = iota
	Clubs
	Hearts
	Spades
)

const (
	Ace Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	suit  Suit
	value Value
}

type Cards [52]Card
type Hand [2]Card

var Suits = [4]Suit{Diamonds, Clubs, Hearts, Spades}
var Values = [13]Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
var Deck Cards

func init() {
	for i, suit := range Suits {
		for j, value := range Values {
			idx := i*13 + j
			Deck[idx] = Card{suit, value}
		}
	}
}

func (card Card) isAce() bool {
	return card.value == 0
}

func (hand Hand) isBusted() bool {
	return hand[0].value+hand[1].value > 21
}

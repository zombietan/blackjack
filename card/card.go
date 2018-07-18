package card

import (
	"fmt"
	"math/rand"
	"time"
)

// suit
const (
	CLUB  = "♧"
	DIA   = "♢"
	HEART = "♡"
	SPADE = "♤"
)

// rank
const (
	ACE = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

//　カード総数
const oneSet = 52

// Card include suit and rank
type Card struct {
	Suit string
	Rank int
}

func (c Card) String() string {
	return fmt.Sprintf("%s%d", c.Suit, c.Rank)
}

// Deck ...
type Deck []Card

var suits = []string{CLUB, DIA, HEART, SPADE}
var ranks = []int{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}

// OneSet Deck initalize
func OneSet() Deck {
	deck := make([]Card, oneSet)
	idx := 0
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(ranks); j++ {
			card := Card{
				Suit: suits[i],
				Rank: ranks[j],
			}
			deck[idx] = card
			idx++
		}
	}
	return deck
}

// Shuffle ...
func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixNano())
	n := len(d)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}

// Pop ...
func (d Deck) Pop() (Card, Deck) {
	card := d[len(d)-1]
	deck := d[:len(d)-1]
	return card, deck
}

// ConvertStringSlice ...
func (d Deck) ConvertStringSlice() []string {
	converted := make([]string, len(d))
	for i := 0; i < len(d); i++ {
		converted[i] = d[i].String()
	}
	return converted
}

package game

import (
	"fmt"
	"math"
	"strings"

	"github.com/zombietan/blackjack/card"
)

// BLACKJACK rank
const BLACKJACK = 21

// Player ...
type Player struct {
	Name string
	Hand []card.Card
}

// Draw ...
func (p Player) Draw(deck card.Deck) (Player, card.Deck) {
	card, drawedDeck := deck.Pop()
	p.Hand = append(p.Hand, card)
	return p, drawedDeck
}

// ShowHand ...
func (p Player) ShowHand() {
	var hand card.Deck = p.Hand
	s := strings.Join(hand.ConvertStringSlice(), " ")
	fmt.Printf("%s's cards: %s\n", p.Name, s)
}

func isContainsAce(cards []card.Card) bool {
	for _, v := range cards {
		if card.ACE == v.Rank {
			return true
		}
	}
	return false
}

// Point ...
func (p Player) Point() int {
	var sum float64
	for _, v := range p.Hand {
		sum += math.Min(float64(v.Rank), 10)
	}
	if (int(sum) <= 11) && (isContainsAce(p.Hand)) {
		return int(sum) + 10
	}
	return int(sum)
}

// Busts ...
func (p Player) Busts() bool {
	return p.Point() > BLACKJACK
}

// Judge ...
func Judge(dealer Player, guest Player) Player {
	if guest.Busts() {
		return dealer
	}
	if dealer.Busts() {
		return guest
	}
	if dealer.Point() < guest.Point() {
		return guest
	}
	return dealer
}

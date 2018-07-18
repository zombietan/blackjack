package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zombietan/blackjack/card"
	"github.com/zombietan/blackjack/game"
)

func main() {
	fmt.Printf("********************\n")
	fmt.Printf("* Blackjack\n")

	var deck card.Deck
	deck = card.OneSet().Shuffle().Shuffle()

	dealer := game.Player{
		Name: "Dealer",
		Hand: []card.Card{},
	}

	guest := game.Player{
		Name: "Guest",
		Hand: []card.Card{},
	}

	// 2枚配る
	for i := 0; i < 2; i++ {
		dealer, deck = dealer.Draw(deck)
		guest, deck = guest.Draw(deck)
	}

	fmt.Printf("\n")
	fmt.Printf("%s's 1st card: %s\n", dealer.Name, dealer.Hand[0])

	fmt.Printf("\n")
	fmt.Printf("* %s's turn *\n", guest.Name)

	guest.ShowHand()
	for {
		if !guest.Busts() {
			decision := decide()
			if decision == STAND {
				break
			}
			if decision == HIT {
				guest, deck = guest.Draw(deck)
				guest.ShowHand()
			}
		} else {
			fmt.Printf("Bust!\n")
			result(dealer, guest)
			return
		}
	}

	fmt.Printf("\n")
	fmt.Printf("* %s's turn *\n", dealer.Name)
	dealer.ShowHand()
	for {
		if !dealer.Busts() {
			if dealer.Point() >= MustDraw {
				break
			} else {
				dealer, deck = dealer.Draw(deck)
				dealer.ShowHand()
			}
		} else {
			fmt.Printf("Bust!\n")
			result(dealer, guest)
			return
		}
	}
	result(dealer, guest)

}

// Choice ...
type Choice int

// choice
const (
	HIT Choice = iota
	STAND
)

// MustDraw Dealer MUST DRAW TO 16 AND STAND ON ALL 17’S
const MustDraw = 17

func (c Choice) String() string {
	switch c {
	case HIT:
		return "h"
	case STAND:
		return "s"
	default:
		return "Unknown"
	}
}

func decide() Choice {
	fmt.Println("Hit/Stand?")
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("[h,s]:")
		sc.Scan()
		s := sc.Text()
		if s == HIT.String() {
			return HIT
		} else if s == STAND.String() {
			return STAND
		} else {
			fmt.Printf("Oops, %s is invalid!\n", s)
		}
	}
}

func result(dealer game.Player, guest game.Player) {
	fmt.Printf("\n")
	fmt.Printf("* Judge *\n")
	winner := game.Judge(dealer, guest)
	fmt.Printf("%s win!", winner.Name)
}

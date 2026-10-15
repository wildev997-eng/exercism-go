package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardList := map[string]int{
		"ace":   11,
		"king":  10,
		"queen": 10,
		"jack":  10,
		"ten":   10,
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
		"one":   1,
	}
	return cardList[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	verdict := ""
	switch {
	case ParseCard(card1)+ParseCard(card2) == 22:
		verdict = "P"
	case ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) < 10:
		verdict = "W"
	case ParseCard(card1)+ParseCard(card2) <= 11:
		verdict = "H"
	case ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) < 7:
		verdict = "S"
	case ParseCard(card1)+ParseCard(card2) >= 17 || ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) >= 10:
		verdict = "S"
	default:
		verdict = "H"
	}
	return verdict
}

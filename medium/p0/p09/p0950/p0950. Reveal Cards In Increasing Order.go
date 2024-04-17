package p0950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	N := len(deck)
	result := make([]int, N)
	skip := false
	indexInDeck := 0
	indexInResult := 0

	sort.Ints(deck)

	for indexInDeck < N {
		// There is an available gap in result
		if result[indexInResult] == 0 {
			// Add a card to result
			if !skip {
				result[indexInResult] = deck[indexInDeck]
				indexInDeck++
			}

			// Toggle skip to alternate between adding and skipping cards
			skip = !skip
		}

		// Progress to the next index of result array
		indexInResult = (indexInResult + 1) % N
	}

	return result
}

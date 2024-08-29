package day_7

import "testing"

func TestDetermineHigherCard(t *testing.T) {

	t.Run("hand1 should be higher than hand2", func(t *testing.T) {
		hand1 := []rune{'A', 'K', '7', '6', '5'}
		hand2 := []rune{'9', '8', '7', '6', '5'}
		hand1Higher := DetermineHigherCard(hand1, hand2, false)

		if hand1Higher != true {
			t.Error("hand1 should be higher than hand2")
		}
	})

	t.Run("hand 1 with first two same should still be higher", func(t *testing.T) {
		hand1 := []rune{'A', 'K', '7', '7', '5'}
		hand2 := []rune{'A', 'K', '7', '6', '5'}
		hand1higher := DetermineHigherCard(hand1, hand2, false)

		if hand1higher != true {
			t.Error("hand1 should be higher than hand2")
		}
	})
}

func TestDetermineHandPower(t *testing.T) {
	testCases := []struct {
		hand          []rune
		expectedPower string
	}{
		{[]rune{'A', 'S', '7', '6', '5'}, "High_card"},
		{[]rune{'2', '5', '3', '4', '2'}, "Pair"},
		{[]rune{'J', 'J', 'Q', 'K', 'Q'}, "Two_pair"},
		{[]rune{'J', 'J', 'Q', 'J', '3'}, "Three_of_a_kind"},
		{[]rune{'J', 'J', 'Q', 'J', 'Q'}, "Full_house"},
		{[]rune{'J', 'J', 'Q', 'J', 'J'}, "Four_of_a_kind"},
		{[]rune{'A', 'A', 'A', 'A', 'A'}, "Five_of_a_kind"},
	}

	for _, tc := range testCases {
		t.Run(tc.expectedPower, func(t *testing.T) {
			handPower := DetermineHandPower(tc.hand, false)
			if handPower != tc.expectedPower {
				t.Errorf("Expected %s, but got %s for hand %v", tc.expectedPower, handPower, tc.hand)
			}
		})
	}
}

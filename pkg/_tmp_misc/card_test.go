package gameobjects

import "testing"

func TestCreateCardRankMap(t *testing.T) {
	Info.Println(ThisFunc())

	crm := createCardRankMap()

	cases := []struct {
		inputProvided  cardRank
		outputExpected int
	}{
		{rA, 14},
		{r2, 2},
	}

	for _, c := range cases {
		// cardIndex := crm[c.inputProvided]

		if crm[c.inputProvided] != c.outputExpected {
			t.Errorf("expected: %v, got %v", c.outputExpected, crm[c.inputProvided])
		}
	}
}

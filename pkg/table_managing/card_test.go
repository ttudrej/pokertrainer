package tableitems

import (
	"testing"
	// "github.com/ttudrej/pokertrainer/pkg/tableitems"
)

func TestCreateCardRankMap(t *testing.T) {
	// Info.Println(ThisFunc())

	crm, _ := CardRankMapCreator.Create(CardRankMapStruct{})

	cases := []struct {
		inputProvided  CardRank
		outputExpected int
	}{
		{RA, 14},
		{R2, 2},
	}

	for _, c := range cases {
		// cardIndex := crm[c.inputProvided]

		if crm[c.inputProvided] != c.outputExpected {
			t.Errorf("expected: %v, got %v", c.outputExpected, crm[c.inputProvided])
		}
	}
}

// EOF

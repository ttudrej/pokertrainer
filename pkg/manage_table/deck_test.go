package manage_table

import "testing"

func TestCreateDeck(t *testing.T) {

	// Create a global card deck, since there will be only one, for any one hand played.
	cdmPtr, olPtr, olfPtr, _ := createDeck()
	cdm := *cdmPtr
	ol := *olPtr
	olf := *olfPtr

	// ############
	// Test the cdm
	casesCdm := []struct {
		inputProvidedCr   cardRank
		inputProvidedCs   cardSuit
		outputExpectedCr  cardRank
		outputExpectedCs  cardSuit
		outputExpectedSeq int
	}{
		{rA, s, rA, s, 1},
	}

	for _, c := range casesCdm {
		if cdm[cdmKey{c.inputProvidedCr, c.inputProvidedCs}].rank != c.outputExpectedCr {
			t.Errorf("cdm; Improper rank assignment, expected: %v, got %v", cdm[cdmKey{c.inputProvidedCr, c.inputProvidedCs}].rank, c.outputExpectedCr)
		}
		if cdm[cdmKey{c.inputProvidedCr, c.inputProvidedCs}].suit != c.outputExpectedCs {
			t.Errorf("cdm; Improper suit assignment, expected: %v, got %v", cdm[cdmKey{c.inputProvidedCr, c.inputProvidedCs}].suit, c.outputExpectedCs)
		}
		if cdm[cdmKey{c.inputProvidedCr, c.inputProvidedCs}].sequence != c.outputExpectedSeq {
			t.Errorf("cdm; Improper sequence assignment, expected: %v, got %v", cdm[cdmKey{c.inputProvidedCr, c.inputProvidedCs}].sequence, c.outputExpectedSeq)
		}
	}

	// #####################
	// Test the Ordered List
	casesOl := []struct {
		inputProvidedSeq  int
		outputExpectedCr  cardRank
		outputExpectedCs  cardSuit
		outputExpectedSeq int
	}{
		{0, rA, s, 1},
		{51, r2, d, 52},
	}

	for _, c := range casesOl {
		if ol[c.inputProvidedSeq].rank != c.outputExpectedCr {
			t.Errorf("ol; Improper rank assignment, expected: %v, got %v", ol[c.inputProvidedSeq].rank, c.outputExpectedCr)
		}
		if ol[c.inputProvidedSeq].suit != c.outputExpectedCs {
			t.Errorf("ol; Improper suit assignment, expected: %v, got %v", ol[c.inputProvidedSeq].suit, c.outputExpectedCs)
		}
		if ol[c.inputProvidedSeq].sequence != c.outputExpectedSeq {
			t.Errorf("ol; Improper sequence assignment, expected: %v, got %v", ol[c.inputProvidedSeq].sequence, c.outputExpectedSeq)
		}
	}

	// ##########################
	// Test the Full Ordered List
	casesOlf := []struct {
		inputProvidedSeq  int
		outputExpectedCr  cardRank
		outputExpectedCs  cardSuit
		outputExpectedSeq int
	}{
		{0, rA, s, 1},
		{51, r2, d, 52},
		{55, rA, d, 4},
	}

	for _, c := range casesOlf {
		if olf[c.inputProvidedSeq].rank != c.outputExpectedCr {
			t.Errorf("olf; Improper rank assignment, expected: %v, got %v", olf[c.inputProvidedSeq].rank, c.outputExpectedCr)
		}
		if olf[c.inputProvidedSeq].suit != c.outputExpectedCs {
			t.Errorf("olf; Improper suit assignment, expected: %v, got %v", olf[c.inputProvidedSeq].suit, c.outputExpectedCs)
		}
		if olf[c.inputProvidedSeq].sequence != c.outputExpectedSeq {
			t.Errorf("olf; Improper sequence assignment, expected: %v, got %v", olf[c.inputProvidedSeq].sequence, c.outputExpectedSeq)
		}
	}
}

// EOF

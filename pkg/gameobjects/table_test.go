/*
File: deck_shuffler_test.go, version history:
v0.1	yyyy/mm/dd	Tomasz Tudrej

*/

package gameobjects

import "testing"

func TestCreateTable(t *testing.T) {

	cases := []struct {
		inputProvided_numberSeats int
		inputProvided_tableID     int
		inputProvided_gameType    gameType
	}{
		{9, 1, nlh},
		// {2, 101, nlh},
	}
	for _, c := range cases {
		// Create a new table
		tPtr, err := createTable(c.inputProvided_numberSeats, c.inputProvided_tableID, c.inputProvided_gameType)
		if err != nil {
			t.Errorf("Could not create a table, seats: %v, type: %v; %v", c.inputProvided_numberSeats, c.inputProvided_gameType, err)
		}
		if tPtr.numberSeats != c.inputProvided_numberSeats {
			t.Errorf("Seats, requested: %v, provided: %v; ", c.inputProvided_numberSeats, tPtr.numberSeats)
		}
	}
}

// EOF

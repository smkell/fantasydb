package main

import (
	"bytes"
	"testing"
)

func TestParseCsv(t *testing.T) {
	csv := `1,9,QB,"Luck, Andrew",IND,10,421,5026,37,17,64,322,3,4,0,0,0,468,$41`
	buf := bytes.NewBufferString(csv)
	expect := Player{
		FirstName: "Andrew",
		LastName:  "Luck",
		Position:  "QB",
		SeasonStats: []SeasonStat{
			SeasonStat{
				Year:      2015,
				Projected: true,
				Team:      "IND",
				ByeWeek:   10,
			},
		},
	}

	actuals, err := ParseCsv(buf, 2015, true)
	if err != nil {
		t.Error("Unexpected error occured:", err)
	}

	if len(actuals) != 1 {
		t.Fatalf("An unexpected number of player records were parsed. Expected %d but found %d", 1, len(actuals))
	}

	actual := actuals[0]

	if expect.FirstName != actual.FirstName {
		t.Errorf("Expected the player's first name to be %s but was %s", expect.FirstName, actual.FirstName)
	}
}

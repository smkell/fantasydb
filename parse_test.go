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
				Year:          2015,
				Projected:     true,
				Team:          "IND",
				ByeWeek:       10,
				ADP:           9,
				Completions:   421,
				PassingYards:  5026,
				PassingTDs:    37,
				Interceptions: 17,
				RushAttempts:  64,
				RushYards:     322,
				RushTDs:       3,
				Fumbles:       4,
				Receptions:    0,
				ReceiveYards:  0,
				ReceiveTDs:    0,
				FantasyPoints: 468,
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

	if expect.LastName != actual.LastName {
		t.Errorf("Expected the player's last name to be %s but was %s", expect.LastName, actual.LastName)
	}

	if expect.Position != actual.Position {
		t.Errorf("Expected the player's position to be %s but was %s", expect.Position, actual.Position)
	}

	if len(actual.SeasonStats) != 1 {
		t.Fatalf("An unexpected number of season stat records were parsed. Expected %d but found %d", 1, len(actual.SeasonStats))
	}

	actualStat := actual.SeasonStats[0]
	expectStat := expect.SeasonStats[0]

	if expectStat.Year != actualStat.Year {
		t.Errorf("Expected the stats to be for year %d but was %d.", expectStat.Year, actualStat.Year)
	}

	if expectStat.Projected != actualStat.Projected {
		t.Errorf("Expected the projection to be %v but was %v", expectStat.Projected, actualStat.Projected)
	}

	if expectStat.Team != actualStat.Team {
		t.Errorf("Expected the player's team to be %s but was %s", expectStat.Team, actualStat.Team)
	}

	if expectStat.ByeWeek != actualStat.ByeWeek {
		t.Errorf("Expected the player to have bye week %d but was %d", expectStat.ByeWeek, actualStat.ByeWeek)
	}

	if expectStat.ADP != actualStat.ADP {
		t.Errorf("Expected the player to have ADP %d but was %d", expectStat.ADP, actualStat.ADP)
	}

	if expectStat.Interceptions != actualStat.Interceptions {
		t.Errorf("Expected the player to have %d Interceptions, but had %d", expectStat.Interceptions, actualStat.Interceptions)
	}

	if expectStat.PassingYards != actualStat.PassingYards {
		t.Errorf("Expected the player to have %d passing yards, but had %d", expectStat.PassingYards, actualStat.PassingYards)
	}

	if expectStat.PassingTDs != actualStat.PassingTDs {
		t.Errorf("Expected the player to have %d passing TDs, but had %d", expectStat.PassingTDs, actualStat.PassingTDs)
	}

	if expectStat.Interceptions != actualStat.Interceptions {
		t.Errorf("Expected the player to have %d interceptions, but had %d", expectStat.Interceptions, actualStat.Interceptions)
	}

	if expectStat.RushAttempts != actualStat.RushAttempts {
		t.Errorf("Expected the player to have %d rushing attempts, but had %d", expectStat.RushAttempts, actualStat.RushAttempts)
	}

	if expectStat.RushTDs != actualStat.RushTDs {
		t.Errorf("Expected the player to have %d rush TDs, but had %d", expectStat.RushTDs, actualStat.RushTDs)
	}

	if expectStat.RushYards != actualStat.RushYards {
		t.Errorf("Expected the player to have %d rush yards, but had %d", expectStat.RushYards, actualStat.RushYards)
	}

	if expectStat.Fumbles != actualStat.Fumbles {
		t.Errorf("Expected the player to have %d fumbles, but had %d", expectStat.Fumbles, actualStat.Fumbles)
	}

	if expectStat.Receptions != actualStat.Receptions {
		t.Errorf("Expected the player to have %d receptions, but had %d", expectStat.Receptions, actualStat.Receptions)
	}

	if expectStat.ReceiveTDs != actualStat.ReceiveTDs {
		t.Errorf("Expected the player to have %d receive TDs, but had %d", expectStat.ReceiveTDs, actualStat.ReceiveTDs)
	}

	if expectStat.ReceiveYards != actualStat.ReceiveYards {
		t.Errorf("Expected the player to have %d receive yards, but had %d", expectStat.ReceiveYards, actualStat.ReceiveYards)
	}

	if expectStat.FantasyPoints != actualStat.FantasyPoints {
		t.Errorf("Expected the player to have %f fantasy points, but had %f", expectStat.FantasyPoints, actualStat.FantasyPoints)
	}
}

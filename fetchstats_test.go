package main

import (
	"testing"
)

func TestFetchStats(t *testing.T) {
	player := Player{
		LastName:  "Luck",
		FirstName: "Andrew",
	}

	expect := []SeasonStat{
		SeasonStat{
			Year:          2012,
			Projected:     false,
			Team:          "IND",
			ByeWeek:       0,
			ADP:           0,
			Completions:   339,
			PassingYards:  4374,
			PassingTDs:    23,
			Interceptions: 18,
			RushAttempts:  62,
			RushYards:     255,
			RushTDs:       5,
			Fumbles:       5,
			Receptions:    0,
			ReceiveYards:  0,
			ReceiveTDs:    0,
			FantasyPoints: 315,
		},
		SeasonStat{
			Year:          2013,
			Projected:     false,
			Team:          "IND",
			ByeWeek:       0,
			ADP:           0,
			Completions:   343,
			PassingYards:  3822,
			PassingTDs:    23,
			Interceptions: 9,
			RushAttempts:  63,
			RushYards:     377,
			RushTDs:       4,
			Fumbles:       2,
			Receptions:    0,
			ReceiveYards:  0,
			ReceiveTDs:    0,
			FantasyPoints: 315,
		},
		SeasonStat{
			Year:          2014,
			Projected:     false,
			Team:          "IND",
			ByeWeek:       0,
			ADP:           0,
			Completions:   380,
			PassingYards:  4761,
			PassingTDs:    40,
			Interceptions: 16,
			RushAttempts:  64,
			RushYards:     273,
			RushTDs:       3,
			Fumbles:       6,
			Receptions:    0,
			ReceiveYards:  0,
			ReceiveTDs:    0,
			FantasyPoints: 315,
		},
	}

	actual, err := FetchStats(player, 0, 2014)

	if err != nil {
		t.Fatal("An unexpected error occured.", err)
	}

	if len(actual) != len(expect) {
		t.Fatalf("An unexpected number of results were fetched. Expected %d, got %d", len(expect), len(actual))
	}

	for i, expectStat := range expect {
		actualStat := actual[i]
		if expectStat.Year != actualStat.Year {
			t.Errorf("Expected stat %d to have year %d but had year %d", i, expectStat.Year, actualStat.Year)
		}

		if expectStat.Team != actualStat.Team {
			t.Errorf("Expected player to be on team %s but was on team %s", expectStat.Team, actualStat.Team)
		}

		if expectStat.Completions != actualStat.Completions {
			t.Errorf("Expected player to have %d completions but had %d", expectStat.Completions, actualStat.Completions)
		}

		if expectStat.PassingYards != actualStat.PassingYards {
			t.Errorf("Expected player to have %d passing yards but had %d", expectStat.PassingYards, actualStat.PassingYards)
		}

		if expectStat.PassingTDs != actualStat.PassingTDs {
			t.Errorf("Expected player to have %d passing TDs but had %d", expectStat.PassingTDs, actualStat.PassingTDs)
		}

		if expectStat.Interceptions != actualStat.Interceptions {
			t.Errorf("Expected player to have %d interceptions but had %d", expectStat.Interceptions, actualStat.Interceptions)
		}

		if expectStat.RushAttempts != actualStat.RushAttempts {
			t.Errorf("Expected player to have %d rushing attempts but had %d", expectStat.RushAttempts, actualStat.RushAttempts)
		}

		if expectStat.RushYards != actualStat.RushYards {
			t.Errorf("Expected player to have %d rushing yards but had %d", expectStat.RushYards, actualStat.RushYards)
		}

		if expectStat.RushTDs != actualStat.RushTDs {
			t.Errorf("Expected player to have %d rushing TDs but had %d", expectStat.RushTDs, actualStat.RushTDs)
		}

		if expectStat.Receptions != actualStat.Receptions {
			t.Errorf("Expected player to have %d receptions but had %d", expectStat.Receptions, actualStat.Receptions)
		}

		if expectStat.ReceiveYards != actualStat.ReceiveYards {
			t.Errorf("Expected player to have %d Receiveing yards but had %d", expectStat.ReceiveYards, actualStat.ReceiveYards)
		}

		if expectStat.ReceiveTDs != actualStat.ReceiveTDs {
			t.Errorf("Expected player to have %d Receiveing TDs but had %d", expectStat.ReceiveTDs, actualStat.ReceiveTDs)
		}

		if expectStat.Fumbles != actualStat.Fumbles {
			t.Errorf("Expected player to have %d fumbles but had %d", expectStat.Fumbles, actualStat.Fumbles)
		}
	}
}

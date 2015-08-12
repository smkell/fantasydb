package main

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

// ParseCsv parses player stats for a given year out of a CSV stream.
func ParseCsv(r io.Reader, year int64, projected bool) ([]Player, error) {
	reader := csv.NewReader(r)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	players := make([]Player, 0, len(records))
	for _, record := range records {
		player, err := parsePlayerStatsFromRecord(record, year, projected)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}

func parsePlayerStatsFromRecord(record []string, year int64, projected bool) (Player, error) {
	const (
		ADP_POS           = 1
		POSITION_POS      = 2
		NAME_POS          = 3
		TEAM_POS          = 4
		BYE_POS           = 5
		COMPLETE_POS      = 6
		PASSING_YARDS_POS = 7
		PASSING_TDS_POS   = 8
		INTERCEPTIONS_POS = 9
		RUSH_ATTEMPTS_POS = 10
		RUSH_YARDS_POS    = 11
		RUSH_TDS_POS      = 12
		FUMBLES_POS       = 13
		RECEPTIONS_POS    = 14
		RECEIVE_YARDS_POS = 15
		RECEIVE_TDS_POS   = 16
		FANTASY_POS       = 17
	)

	var err error

	player := Player{}
	stat := SeasonStat{
		Year:      year,
		Projected: projected,
	}

	stat.ADP, err = strconv.ParseInt(record[ADP_POS], 10, 64)
	if err != nil {
		stat.ADP = 0
	}

	player.Position = record[POSITION_POS]

	fullName := strings.Split(record[NAME_POS], ",")
	player.LastName = strings.Trim(fullName[0], " ")
	player.FirstName = strings.Trim(fullName[1], " ")

	stat.Team = record[TEAM_POS]

	stat.ByeWeek, err = strconv.ParseInt(record[BYE_POS], 10, 64)
	if err != nil {
		stat.ByeWeek = 0
	}

	stat.Completions, err = strconv.ParseInt(record[COMPLETE_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.PassingYards, err = strconv.ParseInt(record[PASSING_YARDS_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.PassingTDs, err = strconv.ParseInt(record[PASSING_TDS_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.Interceptions, err = strconv.ParseInt(record[INTERCEPTIONS_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.RushAttempts, err = strconv.ParseInt(record[RUSH_ATTEMPTS_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.RushTDs, err = strconv.ParseInt(record[RUSH_TDS_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.RushYards, err = strconv.ParseInt(record[RUSH_YARDS_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.Fumbles, err = strconv.ParseInt(record[FUMBLES_POS], 10, 64)
	if err != nil {
		return Player{}, err
	}

	stat.FantasyPoints, err = strconv.ParseFloat(record[FANTASY_POS], 64)
	if err != nil {
		return Player{}, err
	}

	player.SeasonStats = []SeasonStat{stat}
	return player, nil
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchStats(player Player, startYear, endYear int64) ([]SeasonStat, error) {
	const (
		FOOTBALLDB_URL = "http://www.footballdb.com"

		YEAR_POS          = 0
		TEAM_POS          = 1
		PASS_ATTEMPT_POS  = 2
		PASS_COMPLETE_POS = 3
		PASS_YARDS_POS    = 4
		PASS_TDS_POS      = 5
		PASS_INT_POS      = 6
		PASS_2PT_POS      = 7
		RUSH_ATTEMPT_POS  = 8
		RUSH_YARDS_POS    = 9
		RUSH_TDS_POS      = 10
		RUSH_2PT_POS      = 11
		RECEPTIONS_POS    = 12
		REC_YARDS_POS     = 13
		REC_TDS_POS       = 14
		REC_2PT_POS       = 15
		FUMBLES_POS       = 16
	)

	searchName := strings.Split(player.LastName, " ")[0]
	searchUrl := FOOTBALLDB_URL + "/players/players.html?q=" + searchName

	doc, err := goquery.NewDocument(searchUrl)
	if err != nil {
		return nil, err
	}

	fullName := player.LastName + ", " + player.FirstName

	table := doc.Find("table.statistics > tbody > tr > td > a:contains(\"" + fullName + "\")").Parent().Parent().Has("td:contains(\"" + player.Position + "\")")
	s := table.Find("a:contains(\"" + fullName + "\")")
	statUrl, found := s.Attr("href")

	if !found {
		return nil, fmt.Errorf("Unable to find player with name %s", fullName)
	}

	log.Println(FOOTBALLDB_URL + statUrl)
	statDoc, err := goquery.NewDocument(FOOTBALLDB_URL + statUrl)
	if err != nil {
		return nil, err
	}

	headerSel := statDoc.Find("div.divider > h2:contains(\"Fantasy Statistics\")").Parent()
	tableSel := headerSel.Next()

	stats := make([]SeasonStat, 0)
	if startYear < 1960 {
		startYear = 1960
	}

	for curYear := startYear; curYear <= endYear; curYear = curYear + 1 {
		rowSel := tableSel.Find("tbody > tr > td.center:contains(\"" + fmt.Sprint(curYear) + "\")").Parent()

		if rowSel.Size() > 0 {
			stat := SeasonStat{
				Year:      curYear,
				Projected: false,
			}

			rowSel.Find("td").Each(func(i int, s *goquery.Selection) {
				switch i {
				case TEAM_POS:
					stat.Team = strings.ToUpper(s.Text())
				case PASS_COMPLETE_POS:
					stat.Completions, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						log.Println(err)
						stat.Completions = 0
					}
				case PASS_YARDS_POS:
					stat.PassingYards, err = strconv.ParseInt(strings.Replace(s.Text(), ",", "", -1), 10, 64)
					if err != nil {
						log.Println(err)
						stat.PassingYards = 0
					}
				case PASS_TDS_POS:
					stat.PassingTDs, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.PassingTDs = 0
					}
				case PASS_INT_POS:
					stat.Interceptions, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.Interceptions = 0
					}
				case RUSH_ATTEMPT_POS:
					stat.RushAttempts, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.RushAttempts = 0
					}
				case RUSH_TDS_POS:
					stat.RushTDs, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.RushTDs = 0
					}
				case RUSH_YARDS_POS:
					stat.RushYards, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.RushYards = 0
					}
				case RECEPTIONS_POS:
					stat.Receptions, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.Receptions = 0
					}
				case REC_TDS_POS:
					stat.ReceiveTDs, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.ReceiveTDs = 0
					}
				case REC_YARDS_POS:
					stat.ReceiveYards, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.ReceiveYards = 0
					}
				case FUMBLES_POS:
					stat.Fumbles, err = strconv.ParseInt(s.Text(), 10, 64)
					if err != nil {
						stat.Fumbles = 0
					}
				}
			})

			stats = append(stats, stat)
		}
	}
	return stats, nil
}

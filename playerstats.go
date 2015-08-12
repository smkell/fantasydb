package main

// Player represents a fantasy football player.
type Player struct {
	ID          string       `bson:"_id",omitempty`
	FirstName   string       `bson:"first_name"`
	LastName    string       `bson:"last_name"`
	Position    string       `bson:"position"`
	SeasonStats []SeasonStat `bson:"season_stats"`
}

// SeasonStat represents the player's statistics for a particular year.
type SeasonStat struct {
	Year          int64   `bson:"year"`
	Projected     bool    `bson:"projected"`
	Team          string  `bson:"team"`
	ADP           int64   `bson:"adp"`
	ByeWeek       int64   `bson:"bye_week"`
	Completions   int64   `bson:"completions"`
	PassingYards  int64   `bson:"passing_yards"`
	PassingTDs    int64   `bson:"passing_tds"`
	Interceptions int64   `bson:"interceptions"`
	RushAttempts  int64   `bson:"rush_attempts"`
	RushYards     int64   `bson:"rush_yards"`
	RushTDs       int64   `bson:"rush_tds"`
	Fumbles       int64   `bson:"fumbles"`
	Receptions    int64   `bson:"receptions"`
	ReceiveYards  int64   `bson:"receive_yards"`
	ReceiveTDs    int64   `bson:"receive_tds"`
	FantasyPoints float64 `bson:"fantasy_points"`
}

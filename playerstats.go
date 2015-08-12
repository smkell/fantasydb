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
	Year      int64  `bson:"year"`
	Projected bool   `bson:"projected"`
	Team      string `bson:"team"`
}

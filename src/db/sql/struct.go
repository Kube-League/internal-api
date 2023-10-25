package sql

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	gorm.Model
	Name       string       `gorm:"unique;" json:"name"`
	Motto      string       `json:"motto"`
	Logo       string       `gorm:"unique;" json:"logo"`
	Country    string       `json:"country"`
	Town       string       `json:"town"`
	Matches    []Match      `gorm:"many2many:match_teams;" json:"matches"`
	MatchesWon []Match      `gorm:"foreignKey:WinnerID" json:"matches_won"`
	Results    []TeamResult `json:"results"`
	Players    []Player     `json:"players"`
}

type Player struct {
	gorm.Model
	LastName  string              `json:"name"`
	Uuid      string              `gorm:"unique;" json:"uuid"`
	DiscordId string              `gorm:"unique;" json:"discord_id"`
	TeamID    uint                `json:"team"`
	Results   []PlayerMatchResult `json:"results"`
}

type Match struct {
	gorm.Model
	Date          time.Time           `json:"date"`
	WinnerID      uint                `json:"winner"`
	Participants  []Team              `gorm:"many2many:match_teams;" json:"participants"`
	PlayersResult []PlayerMatchResult `json:"players_result"`
	Results       []TeamResult        `gorm:"many2many:match_team_results;" json:"team_results"`
}

type TeamResult struct {
	gorm.Model
	MatchID       uint                `json:"match_id"`
	TeamID        uint                `json:"team_id"`
	Points        uint                `json:"points"`
	PlayersResult []PlayerMatchResult `json:"players_result"`
	Matches       []Match             `gorm:"many2many:match_team_results;" json:"matches"`
}

type PlayerMatchResult struct {
	gorm.Model
	TeamResultID uint `json:"team_result_id"`
	MatchID      uint `json:"match_id"`
	PlayerID     uint `json:"player"`
	Kills        uint `json:"kills"`
	Death        uint `json:"death"`
	PointsScored uint `json:"points_scored"`
}

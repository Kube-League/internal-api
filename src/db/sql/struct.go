package sql

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	gorm.Model
	Name    string
	Motto   string
	Logo    string
	Country string
	Town    string
}

type Player struct {
	gorm.Model
	LastName  string
	Uuid      string
	DiscordId string
	Team      Team
}

type Match struct {
	gorm.Model
	Date         time.Time
	Winner       Team
	Participants []Team        `gorm:"many2many:match_teams;"`
	Results      []MatchResult `gorm:"many2many:match_results;"`
}

type MatchResult struct {
	gorm.Model
	Team          Team
	Points        uint
	PlayersResult []PlayerMatchResult `gorm:"many2many:player_results;"`
}

type PlayerMatchResult struct {
	gorm.Model
	Player       Player
	Kills        uint
	Death        uint
	PointsScored uint
}

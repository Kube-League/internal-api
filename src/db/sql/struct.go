package sql

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	gorm.Model
	Name    string `gorm:"unique;" json:"name"`
	Motto   string `json:"motto"`
	Logo    string `gorm:"unique;" json:"logo"`
	Country string `json:"country"`
	Town    string `json:"town"`
}

type Player struct {
	gorm.Model
	LastName  string `json:"name"`
	Uuid      string `gorm:"unique;" json:"uuid"`
	DiscordId string `gorm:"unique;" json:"discord_id"`
	Team      Team   `json:"team"`
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
	Player       Player `json:"player"`
	Kills        uint   `json:"kills"`
	Death        uint   `json:"death"`
	PointsScored uint   `json:"points_scored"`
}

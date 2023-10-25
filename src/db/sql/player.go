package sql

import (
	"errors"
	"internal-api/src/utils"
)

type PlayerInfo struct {
	Name      string
	Uuid      string
	DiscordId string
	Team      uint
}

var ErrFetchingTeam = errors.New("fetching team in database")

func (p *PlayerInfo) Player() Player {
	var team Team
	err := DB.First(&team, p.Team).Error
	l := utils.Log{Id: "sql.Player"}
	if err != nil {
		l.Error(err)
		return Player{}
	}
	return Player{
		LastName:  p.Name,
		Uuid:      p.Uuid,
		DiscordId: p.DiscordId,
		Team:      team,
	}
}

func CreatePlayer(info PlayerInfo) (uint, error) {
	p := info.Player()
	l := utils.Log{Id: "sql.CreatePlayer"}
	if p.LastName == "" {
		return 0, ErrFetchingTeam
	}
	err := DB.Create(&p).Error
	if err != nil {
		l.Error(err)
		return 0, err
	}
	return p.ID, nil
}

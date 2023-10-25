package sql

import (
	"internal-api/src/utils"
	"time"
)

type MatchInfo struct {
	Date         int64
	Participants []uint
}

func (m *MatchInfo) Match() (Match, error) {
	var parts []Team
	for _, p := range m.Participants {
		var team Team
		err := DB.First(&team, p).Error
		l := utils.Log{Id: "sql.Match"}
		if err != nil {
			l.Error(err)
			return Match{}, ErrFetchingTeam
		}
		parts = append(parts, team)
	}
	return Match{
		Date:         time.Time{},
		Participants: parts,
	}, nil
}

func CreateMatch(info MatchInfo) (uint, error) {
	p, err := info.Match()
	l := utils.Log{Id: "sql.CreateMatch"}
	if err != nil {
		l.Error(err)
		return 0, err
	}
	err = DB.Create(&p).Error
	if err != nil {
		l.Error(err)
		return 0, err
	}
	return p.ID, nil
}

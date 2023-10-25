package sql

import (
	"internal-api/src/utils"
	"time"
)

type MatchInfo struct {
	Date         int64  `json:"date"`
	Participants []uint `json:"participants"`
}

type ResultInfo struct {
	MatchId     uint `json:"match_id"`
	TeamResults []struct {
		TeamID        uint `json:"team_id"`
		Points        uint `json:"points"`
		PlayerResults []struct {
			PlayerID     uint `json:"player_id"`
			Kills        uint `json:"kills"`
			Death        uint `json:"death"`
			PointsScored uint `json:"points_scored"`
		} `json:"player_results"`
	} `json:"team_results"`
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

func (r *ResultInfo) Result() []TeamResult {
	var results []TeamResult
	for _, res := range r.TeamResults {
		var pResults []PlayerMatchResult
		for _, pRes := range res.PlayerResults {
			pResults = append(pResults, PlayerMatchResult{
				MatchID:      r.MatchId,
				PlayerID:     pRes.PlayerID,
				Kills:        pRes.Kills,
				Death:        pRes.Death,
				PointsScored: pRes.PointsScored,
			})
		}
		results = append(results, TeamResult{
			MatchID:       r.MatchId,
			TeamID:        res.TeamID,
			Points:        res.Points,
			PlayersResult: pResults,
		})
	}
	return results
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

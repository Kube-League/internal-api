package sql

import (
	"internal-api/src/utils"
)

type TeamInfo struct {
	Name    string
	Motto   string
	Logo    string
	Country string
	Town    string
}

func (t *TeamInfo) Team() Team {
	return Team{
		Name:    t.Name,
		Motto:   t.Motto,
		Logo:    t.Logo,
		Country: t.Country,
		Town:    t.Town,
	}
}

func CreateTeam(info TeamInfo) (uint, error) {
	t := info.Team()
	l := utils.Log{Id: "sql.CreateTeam"}
	err := DB.Create(&t).Error
	if err != nil {
		l.Error(err)
		return 0, err
	}
	return t.ID, nil
}

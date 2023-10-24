package sql

import "internal-api/src/utils"

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

func CreateTeam(info TeamInfo) uint {
	t := info.Team()
	r := DB.Create(&t)
	if r.Error != nil {
		utils.LogError("sql.CreateTeam", r.Error)
		return 0
	}
	return t.ID
}

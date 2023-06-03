package event

import (
	"dcs/team"
)

type Event struct{
	Name string
	Teams []team.Team
	Date string 
}

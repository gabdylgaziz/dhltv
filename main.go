package main

import (
	"dcs/handlers"
	"dcs/models"
)

var players []models.Player

func main() {
	handlers.HandleRequests()
	//fmt.Println(fmt.Sprintf("KD ratio of %s: %.2f", gg.Nickname, gg.GetKD()))
	//fmt.Println(fmt.Sprintf("Impact of %s: %.2f", gg.Nickname, gg.GetImpact()))
	//fmt.Println(fmt.Sprintf("Rating 2.0 of %s: %.2f", gg.Nickname, gg.GetRating()))
}
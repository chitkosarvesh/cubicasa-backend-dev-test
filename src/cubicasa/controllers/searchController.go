package controllers

import (
	"cubicasa/database"

	"github.com/gin-gonic/gin"
)

//SearchResults the struct that returns the search results
type SearchResults struct {
	Hubs  []database.Hub  `json:"hubs"`
	Teams []database.Team `json:"teams"`
	Users []database.User `json:"users"`
}

func search(c *gin.Context) {
	var hubs []database.Hub
	var teams []database.Team
	var users []database.User
	if len(c.Param("query")) != 0 {
		queryString := "%" + c.Param("query") + "%"
		database.DB.Where("name LIKE ? or location like ?", queryString, queryString).Find(&hubs)
		database.DB.Where("name LIKE ?", queryString).Find(&teams)
		database.DB.Where("first_name LIKE ? or last_name like ? or contact like ? or email like ? or username like ?", queryString, queryString, queryString, queryString, queryString).Find(&users)
		sr := SearchResults{
			Hubs:  hubs,
			Teams: teams,
			Users: users,
		}
		c.JSON(200, &gin.H{
			"success": "true",
			"rows":    sr,
		})
	} else {
		c.JSON(200, &gin.H{
			"success": false,
			"error":   "Search query is blank",
		})
	}
}

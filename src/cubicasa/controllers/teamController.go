package controllers

import (
	"cubicasa/database"

	"github.com/gin-gonic/gin"
)

func teamCreate(c *gin.Context) {
	var team database.Team
	var hub database.Hub
	c.BindJSON(&team)
	result := database.DB.Create(&team)
	if result.Error == nil {
		result1 := database.DB.First(&hub, team.HubID)
		if result1.Error == nil {
			hub.Teams = append(hub.Teams, team)
			result2 := database.DB.Save(&hub)
			if result2.Error == nil {
				c.JSON(200, &gin.H{
					"success": true,
					"data": &gin.H{
						"id":           team.ID,
						"rowsAffected": result.RowsAffected,
					},
				})
			} else {
				c.JSON(403, &gin.H{
					"success": false,
					"data": &gin.H{
						"error": result2.Error.Error(),
					},
				})
			}
		} else {
			c.JSON(403, &gin.H{
				"success": false,
				"data": &gin.H{
					"error": result1.Error.Error(),
				},
			})
		}
	} else {
		c.JSON(403, &gin.H{
			"success": false,
			"data": &gin.H{
				"error": result.Error.Error(),
			},
		})
	}
}
func teamUpdate(c *gin.Context) {}
func teamList(c *gin.Context) {
	var teams []database.Team
	result := database.DB.Find(&teams)
	if result.Error == nil {
		for index, row := range teams {
			database.DB.Model(row).Association("Users").Find(&row.Users)
			teams[index] = row
		}
		c.JSON(200, &gin.H{
			"success": true,
			"rows":    teams,
		})
	} else {
		c.JSON(403, &gin.H{
			"success": false,
			"data": &gin.H{
				"error": result.Error.Error(),
			},
		})
	}
}
func teamReadOne(c *gin.Context) {
	var team database.Team
	result := database.DB.First(&team, c.Param("id"))
	if result.Error == nil {
		database.DB.Model(team).Association("Users").Find(&team.Users)
		c.JSON(200, &gin.H{
			"success": true,
			"data":    team,
		})
	} else {
		c.JSON(403, &gin.H{
			"success": false,
			"data": &gin.H{
				"error": result.Error.Error(),
			},
		})
	}
}
func teamDelete(c *gin.Context) {
	var team database.Team
	result := database.DB.Delete(&team, c.Param("id"))
	if result.Error == nil {
		c.JSON(200, &gin.H{
			"success": true,
			"data": &gin.H{
				"id":           team.ID,
				"rowsAffected": result.RowsAffected,
			},
		})
	} else {
		c.JSON(403, &gin.H{
			"success": false,
			"data": &gin.H{
				"error": result.Error.Error(),
			},
		})
	}
}

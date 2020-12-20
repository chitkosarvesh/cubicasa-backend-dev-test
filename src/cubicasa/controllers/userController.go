package controllers

import (
	"cubicasa/database"

	"github.com/gin-gonic/gin"
)

func userCreate(c *gin.Context) {
	var user database.User
	var team database.Team
	c.BindJSON(&user)
	result := database.DB.Create(&user)
	if result.Error == nil {
		result1 := database.DB.First(&team, user.TeamID)
		if result1.Error == nil {
			team.Users = append(team.Users, user)
			result2 := database.DB.Save(&team)
			if result2.Error == nil {
				c.JSON(200, &gin.H{
					"success": true,
					"data": &gin.H{
						"id":           user.ID,
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
func userUpdate(c *gin.Context) {}
func userList(c *gin.Context) {
	var users []database.User
	result := database.DB.Find(&users)
	if result.Error == nil {
		c.JSON(200, &gin.H{
			"success": true,
			"rows":    users,
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
func userReadOne(c *gin.Context) {
	var user database.User
	result := database.DB.First(&user, c.Param("id"))
	if result.Error == nil {
		c.JSON(200, &gin.H{
			"success": true,
			"data":    user,
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
func userDelete(c *gin.Context) {
	var user database.User
	result := database.DB.Delete(&user, c.Param("id"))
	if result.Error == nil {
		c.JSON(200, &gin.H{
			"success": true,
			"data": &gin.H{
				"id":           user.ID,
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

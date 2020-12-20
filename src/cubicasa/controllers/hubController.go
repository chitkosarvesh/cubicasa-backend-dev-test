package controllers

import (
	"cubicasa/database"

	"github.com/gin-gonic/gin"
)

func hubCreate(c *gin.Context) {
	var hub database.Hub
	c.BindJSON(&hub)
	result := database.DB.Create(&hub)
	if result.Error == nil {
		c.JSON(200, &gin.H{
			"success": true,
			"data": &gin.H{
				"id":           hub.ID,
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
func hubUpdate(c *gin.Context) {}
func hubList(c *gin.Context) {
	var hubs []database.Hub
	result := database.DB.Find(&hubs)
	if result.Error == nil {
		for index, row := range hubs {
			database.DB.Model(row).Association("Teams").Find(&row.Teams)
			hubs[index] = row
		}
		c.JSON(200, &gin.H{
			"success": true,
			"rows":    hubs,
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
func hubReadOne(c *gin.Context) {
	var hub database.Hub
	result := database.DB.First(&hub, c.Param("id"))
	if result.Error == nil {
		database.DB.Model(hub).Association("Teams").Find(&hub.Teams)
		c.JSON(200, &gin.H{
			"success": true,
			"data":    hub,
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
func hubDelete(c *gin.Context) {
	var hub database.Hub
	result := database.DB.Delete(&hub, c.Param("id"))
	if result.Error == nil {
		c.JSON(200, &gin.H{
			"success": true,
			"data": &gin.H{
				"id":           hub.ID,
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

package main

import (
	"cubicasa/config"
	"cubicasa/controllers"
	"cubicasa/database"
)

func main() {
	if ok := config.Setup(); ok {
		database.Setup()
		controllers.Setup()
	}
}

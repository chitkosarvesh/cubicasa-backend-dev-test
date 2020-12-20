package database

import "gorm.io/gorm"

var models []interface{}

type Hub struct {
	gorm.Model
	Name     string `json:"name"`
	Location string `json:"location"`
	Teams    []Team `json:"teams"`
}
type Team struct {
	gorm.Model
	Name  string `json:"name"`
	HubID uint   `json:"hub_id" gorm:"foreignKey:HubID"`
	Users []User `json:"users"`
}
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	TeamID    uint   `json:"team_id" gorm:"foreignKey:TeamID"`
}

func init() {
	models = append(models, &User{})
	models = append(models, &Team{})
	models = append(models, &Hub{})
}

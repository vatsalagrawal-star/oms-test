package models

type User struct {
	AbstractCreateUpdateModel
	Username 					string 	`json:"username" binding:"required" gorm:"unique"`
	Email						string	`json:"email" binding:"required" gorm:"unique"`
	Name						string	`json:"name" binding:"required"`
}


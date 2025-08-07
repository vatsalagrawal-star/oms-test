package models

import "time"

type AbstractCreateUpdateModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint       `gorm:"primarykey" json:"ID"`
	Username    string     `json:"Username`
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"DueDate,omitempty"`
	CreatedAt   *time.Time `json:"CreatedAt"`
	UpdatedAt   *time.Time `json:"UpdatedAt"`
}

type User struct {
	gorm.Model
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

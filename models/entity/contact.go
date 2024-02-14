package entity

import (
	"time"
)

type Contact struct {
	ID        uint      `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Email     string    `gorm:"column:email" json:"email"`
	Handphone string    `gorm:"column:no_hp" json:"no_hp"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
}

// buat table namenya
func (c Contact) TableName() string {
	return "contacts"
}

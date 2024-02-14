package entity

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	UserID    string     `gorm:"column:user_id" json:"user_id"`
	Name      string     `gorm:"column:name" json:"name"`
	Password  string     `gorm:"column:password" json:"-"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
	Contacts  []Contact  `gorm:"foreignKey:user_id;references:id" json:"contacts"`
	// Contacts []Contact
	// Posts     []Post     `gorm:"foreignKey:user_id;references:id" json:"posts"`
}

func (u User) TableName() string {
	return "users"
}

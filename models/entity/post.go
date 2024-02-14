package entity

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(100)" json:"title"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	UserID      int       `gorm:"column:user_id" json:"user_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
}

// buat table namenya
func (p Post) TableName() string {
	return "posts"
}

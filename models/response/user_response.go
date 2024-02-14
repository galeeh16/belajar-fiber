package response

import (
	"galih/belajar-fiber/models/entity"
	"time"
)

type UserResponse struct {
	ID        int              `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	UserID    string           `gorm:"column:user_id" json:"user_id"`
	Name      string           `json:"name"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Contacts  []entity.Contact `json:"contacts" gorm:"foreignKey:user_id;references:id"`
	// Posts     []entity.Post    `json:"posts" gorm:"foreignKey:user_id;references:id"`
}

func (u UserResponse) TableName() string {
	return "users"
}

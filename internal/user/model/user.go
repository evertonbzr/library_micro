package model

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FullName  string    `gorm:"type:varchar(255)" json:"full_name"`
	Email     string    `gorm:"column:email;type:varchar(255);unique" json:"email"`
	Password  string    `gorm:"column:password;type:varchar(255)" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

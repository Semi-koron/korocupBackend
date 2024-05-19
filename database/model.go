package database

import (
	"time"
)

type User struct {
	ID          int    `gorm:"primaryKey"`
	UserName    string `gorm:"unique;not null"`
	FirebaseUID string `gorm:"unique;not null"`
	Icon        string
	Profile     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Post struct {
	ID        int    `gorm:"primaryKey"`
	UserName  string `gorm:"not null"`
	Image     string
	Reply     int
	Likes     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Like struct {
	ID        int    `gorm:"primaryKey"`
	UserName  string `gorm:"not null"`
	PostID    int    `gorm:"not null"`
	CreatedAt time.Time
}

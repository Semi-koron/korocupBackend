package database

import (
	"time"
)

type User struct {
	ID          int `gorm:"primaryKey"`
	UserName    string
	FirebaseUID string
	Icon        int
	Profile     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Post struct {
	ID        int `gorm:"primaryKey"`
	UserName  string
	Image     string
	Reply     int
	Likes     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Like struct {
	ID        int `gorm:"primaryKey"`
	UserName  string
	PostID    int
	CreatedAt time.Time
}

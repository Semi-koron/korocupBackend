package database

import (
	"time"
)

type User struct {
	ID          int `gorm:"primaryKey"`
	UserID      string
	FirebaseUID string
	Icon        int
	Profile     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Post struct {
	ID        int `gorm:"primaryKey"`
	UserID    string
	User      *User
	Image     string
	Reply     int
	Likes     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Like struct {
	ID        int `gorm:"primaryKey"`
	UserID    string
	User      *User
	PostID    int
	Post      *Post
	CreatedAt time.Time
}

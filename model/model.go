package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Comic struct {
	*gorm.Model
	Author      string `json:"author" form:"author"`
	Description string `json:"description" form:"description"`
	Genre       string `json:"genre" form:"genre"`
	Title       string `json:"title" form:"title"`
}

type Chapter struct {
	*gorm.Model
	Title         string  `json:"title" form:"title"`
	Chapternumber string  `json:"chapternumber" form:"chapternumber" `
	Content       string  `json:"content" form:"content"`
	ComicID       []Comic `json:"comicid" form:"comicid"`
}

type Subscribe struct {
	*gorm.Model
	UserID    []User    `json:"userid" form:"userid"`
	ComicID   []Comic   `json:"comicid" form:"comicid"`
	ChapterID []Chapter `json:"chapterid" form:"chapterid"`
}

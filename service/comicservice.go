package service

import (
	"model/model"

	"gorm.io/gorm"
)

type ComicService struct {
	db *gorm.DB
	// Implementasi sesuai kebutuhan
}

func NewComicService(db *gorm.DB) *ComicService {
	return &ComicService{
		db: db,
	}
}

func (s *ComicService) AddComic(comic model.Comic) {
	// Implementasi penambahan komik
}

func (s *ComicService) EditComic(comic model.Comic) {
	// Implementasi pengeditan komik
}

func (s *ComicService) DeleteComic(comicID uint) {
	// Implementasi penghapusan komik
}

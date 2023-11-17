package service

import "model/model"

type ReaderService struct {
	// Implementasi sesuai kebutuhan
}

func (s *ReaderService) GetComicByID(comicID uint) (*model.Comic, error) {
	// Implementasi pengambilan komik berdasarkan ID
}

func (s *ReaderService) GetChapterByNumber(comicID, chapterNumber uint) (*model.Chapter, error) {
	// Implementasi pengambilan bab komik berdasarkan nomor bab
}

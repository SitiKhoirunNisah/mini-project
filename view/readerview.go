package view

import "model/model"

type ReaderView struct {
	// Implementasi sesuai kebutuhan
}

func (v *ReaderView) ShowComicNotFoundError(message string) {
	// Implementasi tampilan pesan kesalahan komik tidak ditemukan
}

func (v *ReaderView) ShowChapterNotFoundError(message string) {
	// Implementasi tampilan pesan kesalahan bab komik tidak ditemukan
}

func (v *ReaderView) DisplayChapter(comic model.Comic, chapter model.Chapter) {
	// Implementasi tampilan bab komik
}

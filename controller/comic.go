package controller

import (
	"model/model"
	"model/service"
	"model/view"
)

type ComicController struct {
	comicService service.ComicService
	comicView    view.ComicView
}

func NewComicController(comicService service.ComicService, comicView view.ComicView) *ComicController {
	return &ComicController{
		comicService: comicService,
		comicView:    comicView,
	}
}

func (c *ComicController) AddComic(author, description, genre, title string) {
	newComic := model.Comic{
		Author:      author,
		Description: description,
		Genre:       genre,
		Title:       title,
	}
	c.comicService.AddComic(newComic)
	c.comicView.ShowComicAddedMessage()
}

func (c *ComicController) EditComic(comicID uint, author, description, genre, title string) {
	updatedComic := model.Comic{
		ID:          comicID,
		Author:      author,
		Description: description,
		Genre:       genre,
		Title:       title,
	}
	c.comicService.EditComic(updatedComic)
	c.comicView.ShowComicEditedMessage()
}

func (c *ComicController) DeleteComic(comicID uint) {
	c.comicService.DeleteComic(comicID)
	c.comicView.ShowComicDeletedMessage()
}

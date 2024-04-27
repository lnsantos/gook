package books

import "gobook/entity"

type ISearchBooksRepository interface {
	Filter(title string) ([]entity.BookDomain, error)
	FindByID(id int32) (entity.BookDomain, error)
}

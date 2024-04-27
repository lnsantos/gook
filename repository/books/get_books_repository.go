package books

import "gobook/entity"

type IGetBooksRepository interface {
	Get() ([]entity.BookDomain, error)
}

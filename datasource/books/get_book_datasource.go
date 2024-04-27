package books

import "gobook/entity"
import "gobook/external/mockdb"

func Get() ([]entity.BookDomain, error) {
	db := mockdb.Start("books")
	r, _ := db.Query("SELECT * FROM books")
	return r.([]entity.BookDomain), nil
}

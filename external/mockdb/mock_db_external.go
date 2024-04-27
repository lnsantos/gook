package mockdb

import "gobook/entity"

type MockDB struct {
	Query func(query string) (any, error)
}

func Start(context string) MockDB {
	switch context {
	case "books":
		return newBookMockImpl()
	}

	return newBookMockImpl()
}

func newBookMockImpl() MockDB {
	return MockDB{
		Query: func(query string) (any, error) {
			return []entity.BookDomain{
				{
					ID:     1,
					Title:  "Title 1",
					Author: "Author 1",
				},
				{
					ID:     2,
					Title:  "Title 2",
					Author: "Author 2",
				},
				{
					ID:     3,
					Title:  "Title 3",
					Author: "Author 3",
				},
			}, nil
		},
	}
}

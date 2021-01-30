package controllers

import (
	"github.com/revel/revel"
	"myapp/app"
)

type App struct {
	*revel.Controller
}

type Book struct {
	Author    string
	BookName  string
	Publisher string
}

func (c App) Index() revel.Result {
	r, err := app.DB.Query("select (book_name, author_fio, publisher_name) "+
		"from book "+
		"where book_name like '$1'",
		"123",
	)
	if err != nil {
		return c.RenderError(err)
	}

	books := make([]Book, 0)

	for r.Next() {
		b := Book{}
		err := r.Scan(&b.BookName, &b.Author, &b.Publisher)
		if err != nil {
			return c.RenderError(err)
		}

		books = append(books, b)
	}
	return c.Render()
}

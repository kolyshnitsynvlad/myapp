package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	type Book struct {
		Author    string
		BookName  string
		Publisher string
	}
	//data->>'id', data->>'type', data->>'title' (book_name, author_fio, publisher_name)
	r, err := DB.Query("select book_name, author_fio, publisher_name " +
		"from book " +
		"join book_author using (book_id) " +
		"left join author using (author_id) " +
		"join book_publisher using (book_id) " +
		"left join publisher using (publisher_id)",
	)
	if err != nil {
		return c.RenderError(err)
	}

	books := make([]Book, 0)

	for r.Next() {
		b := Book{}
		err := r.Scan(&b.BookName, &b.Publisher, &b.Author)
		if err != nil {
			return c.RenderError(err)
		}

		books = append(books, b)
	}

	return c.Render(books)

}

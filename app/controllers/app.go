package controllers

import (
	"github.com/revel/revel"
	"strings"
)

type App struct {
	*revel.Controller
}

func (c App) Index(Author string, BookName string, Publisher string) revel.Result {
	Author = strings.TrimSpace(Author)
	BookName = strings.TrimSpace(BookName)
	Publisher = strings.TrimSpace(Publisher)

	c.Log.Debug(Author)
	c.Log.Debug(BookName)
	c.Log.Debug(Publisher)

	type Book struct {
		Author    string
		BookName  string
		Publisher string
	}
	//data->>'id', data->>'type', data->>'title' (book_name, author_fio, publisher_name)
	strselect := "select book_name, author_fio, publisher_name " +
		"from book " +
		"left join book_author using (book_id) " +
		"left join author using (author_id) " +
		"left join book_publisher using (book_id) " +
		"left join publisher using (publisher_id) "
	var search string
	if len(Author) > 0 {
		if strings.ToLower(Author) == "нет автора" {
			search = "where author_fio is null "
		} else {
			search = "where author_fio ilike '%" + Author + "%' "
		}
	}
	if len(BookName) > 0 {
		if len(search) > 0 {
			search += "and book_name ilike '%" + BookName + "%' "
		} else {
			search = "where book_name ilike '%" + BookName + "%' "
		}
	}
	if len(Publisher) > 0 {
		if len(search) > 0 {
			search += "and publisher_name "
		} else {
			search = "where publisher_name "
		}
		if strings.ToLower(Publisher) == "нет издателя" {
			search += "is null "
		} else {
			search += "ilike '%" + Publisher + "%' "
		}
	}
	strselect += search
	c.Log.Debug(strselect)
	r, err := DB.Query(strselect)
	if err != nil {
		return c.RenderError(err)
	}

	books := make([]Book, 0)

	for r.Next() {
		bi := struct {
			BookName  interface{}
			Publisher interface{}
			Author    interface{}
		}{}
		err := r.Scan(&bi.BookName, &bi.Author, &bi.Publisher)
		if err != nil {
			return c.RenderError(err)
		}
		b := Book{}
		if str, ok := bi.BookName.(string); ok {
			b.BookName = str
		}
		if str, ok := bi.Publisher.(string); ok {
			b.Publisher = str
		}
		if str, ok := bi.Author.(string); ok {
			b.Author = str
		}

		var exists bool
		for i, v := range books {
			if v.BookName == b.BookName && v.Publisher == b.Publisher {
				books[i].Author += ", " + b.Author
				exists = true
				break
			}
		}

		if !exists {
			books = append(books, b)
		}
	}

	return c.Render(books, Author, BookName, Publisher)

}

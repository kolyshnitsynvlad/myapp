package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index(Author string, BookName string, Publisher string) revel.Result {
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
		search = "where author_fio ilike '%" + Author + "%' "
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
			search += "and publisher_name ilike '%" + Publisher + "%' "
		} else {
			search = "where publisher_name ilike '%" + Publisher + "%' "
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
		b := Book{}
		err := r.Scan(&b.BookName, &b.Publisher, &b.Author)
		if b.BookName != "" || b.Publisher != "" || b.Author != "" {
			books = append(books, b)
		} else {
			if err != nil {
				return c.RenderError(err)
			}
		}
		//if err != nil {
		//	return c.RenderError(err)
		//}

		//books = append(books, b)
	}

	return c.Render(books, Author, BookName, Publisher)

}

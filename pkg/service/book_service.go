package service

import(
	userpb "github.com/anjush-bhargavan/golib_user/pkg/pb"
	book "github.com/anjush-bhargavan/golib_user/pkg/books/handler"
)


func (u *UserService) FetchBookByIDService(p *userpb.UBookID) (*userpb.UBookModel,error) {
	result, err := book.FetchBookByIDHandler(u.client,p)
	if err != nil {
		return nil, err
	}

	return &userpb.UBookModel{
		BookName: result.BookName,
		Author: result.Author,
		NoOfCopies: result.NoOfCopies,
		Year: result.Year,
		Publications: result.Publications,
		Category: result.Category,
		Description: result.Description,

	},nil
}

func (u *UserService) FetchBookByNameService(p *userpb.UBookName) (*userpb.UBookModel,error) {
	result, err := book.FetchBookByNameHandler(u.client,p)
	if err != nil {
		return nil, err
	}

	return &userpb.UBookModel{
		BookName: result.BookName,
		Author: result.Author,
		NoOfCopies: result.NoOfCopies,
		Year: result.Year,
		Publications: result.Publications,
		Category: result.Category,
		Description: result.Description,

	},nil
}

func (u *UserService) FetchAllBooksService(p *userpb.UNoParam) (*userpb.UBookList,error) {
	result, err := book.FetchAllBooksHandler(u.client,p)
	if err != nil {
		return nil, err
	}
	var books []*userpb.UBookModel
	for _,i := range result.Books {
		books = append(books, &userpb.UBookModel{
			BookName: i.BookName,
			Author: i.Author,
			NoOfCopies: i.NoOfCopies,
			Year: i.Year,
			Publications: i.Publications,
			Category: i.Category,
			Description: i.Description,
		})
	}

	return &userpb.UBookList{
		Books: books,
	},nil
}
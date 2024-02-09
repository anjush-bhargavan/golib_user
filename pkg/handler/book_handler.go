package handler

import (
	"context"
	"log"

	userpb "github.com/anjush-bhargavan/golib_user/pkg/pb"
)



func (u *UserHandler) UserFetchBookByID(ctx context.Context,p *userpb.UBookID) (*userpb.UBookModel, error) {
	result, err := u.svc.FetchBookByIDService(p)
	if err != nil {
		log.Println("error while fetching book by id")
		return nil, err
	}
	return result, nil
}


func (u *UserHandler) UserFetchBookByName(ctx context.Context, p *userpb.UBookName) (*userpb.UBookModel, error) {
	result, err := u.svc.FetchBookByNameService(p)
	if err != nil {
		log.Println("error while fetching book by name")
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) USerFetchAllBooks(ctx context.Context,p *userpb.UNoParam) (*userpb.UBookList, error) {
	result, err := u.svc.FetchAllBooksService(p)
	if err != nil {
		log.Println("error while fetching all books")
		return nil, err
	}
	return result, nil
}
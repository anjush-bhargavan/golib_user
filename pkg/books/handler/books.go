package handler

import (
	"context"
	"log"

	bookpb "github.com/anjush-bhargavan/golib_user/pkg/books/pb"
	userpb "github.com/anjush-bhargavan/golib_user/pkg/pb"
)

func FetchBookByIDHandler(client bookpb.BookServicesClient, p *userpb.UBookID) (*bookpb.BookModel, error) {
	ctx := context.Background()

	response, err := client.FetchBookByID(ctx, &bookpb.BookID{Id: p.Id})
	if err != nil {
		log.Printf("error while fetching book by name")
		return nil, err
	}
	return response, nil
}

func FetchBookByNameHandler(client bookpb.BookServicesClient, p *userpb.UBookName) (*bookpb.BookModel, error) {
	ctx := context.Background()

	response, err := client.FetchBookByName(ctx, &bookpb.BookName{Name: p.Name})
	if err != nil {
		log.Printf("error while fetching book by id")
		return nil, err
	}
	return response, nil
}

func FetchAllBooksHandler(client bookpb.BookServicesClient, p *userpb.UNoParam) (*bookpb.BookList, error) {
	ctx := context.Background()

	response, err := client.FetchAllBooks(ctx, &bookpb.NoParam{})
	if err != nil {
		log.Printf("error while fetching books")
		return nil, err
	}
	return response, nil
}

package books

import (
	"log"
	bookpb "github.com/anjush-bhargavan/golib_user/pkg/books/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (bookpb.BookServicesClient, error) {
	grpc, err := grpc.Dial(":8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error dialing to grpc client :8083")
		return nil, err
	}
	log.Printf("successfully connected to client :8083")
	return bookpb.NewBookServicesClient(grpc), nil
}

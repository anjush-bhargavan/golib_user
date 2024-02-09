package di

import (
	"log"

	"github.com/anjush-bhargavan/golib_user/config"
	"github.com/anjush-bhargavan/golib_user/pkg/books"
	"github.com/anjush-bhargavan/golib_user/pkg/db"
	"github.com/anjush-bhargavan/golib_user/pkg/handler"
	"github.com/anjush-bhargavan/golib_user/pkg/repo"
	"github.com/anjush-bhargavan/golib_user/pkg/server"
	"github.com/anjush-bhargavan/golib_user/pkg/service"
)

// Init function load the configurations and initialize all instances
func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	client, err := books.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}
	userRepo := repo.NewUserRepo(db)
	userSvc := service.NewUserService(userRepo, client)
	userHandle := handler.NewUserHandler(userSvc)
	server.NewGrpcServer(userHandle)

}

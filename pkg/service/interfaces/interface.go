package service_interfaces

import (
	// dom "github.com/anjush-bhargavan/goilb_user/pkg/DOM"
	pb "github.com/anjush-bhargavan/golib_user/pkg/pb"
)


type UserServiceInter interface {
	SignupService(userpb *pb.SignupRequest) (*pb.CommonResponse, error)
	LoginService(userpb *pb.LoginRequest) (*pb.CommonResponse, error)
	ProfileService(userpb *pb.UserID) (*pb.Profile,error)
	UpdateProfileService(userpb *pb.ProfileUpdate) (*pb.Profile,error)


	FetchBookByIDService(p *pb.UBookID) (*pb.UBookModel,error)
	FetchBookByNameService(p *pb.UBookName) (*pb.UBookModel,error)
	FetchAllBooksService(p *pb.UNoParam) (*pb.UBookList,error)

	FindAllUsersService(p *pb.UNoParam) (*pb.Users, error)
	EditUserService(p *pb.SignupRequest) (*pb.CommonResponse,error) 
	FindUserByEmailService(p *pb.Email) (*pb.Profile,error)
	FindUserByIDService(p *pb.UserID) (*pb.Profile,error)
	DeleteUserService(p *pb.UserID) (*pb.CommonResponse,error)


	FindMembershipByIDService(p *pb.UserID) (*pb.Membership, error)
	CreateMembershipService(p *pb.Membership) (*pb.CommonResponse, error)
	UpdateMembershipService(p *pb.Membership) (*pb.CommonResponse,error)
}

package handler

import (
	"context"
	"fmt"

	pb "github.com/anjush-bhargavan/golib_user/pkg/pb"
)

func (u *UserHandler) FindAllUsers(ctx context.Context, p *pb.UNoParam) (*pb.Users, error) {
	result, err := u.svc.FindAllUsersService(p)
	if err != nil {
		fmt.Println("hii")
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) EditUser(ctx context.Context, p *pb.SignupRequest) (*pb.CommonResponse, error) {
	result, err := u.svc.EditUserService(p)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (u *UserHandler) DeleteUser(ctx context.Context, p *pb.UserID) (*pb.CommonResponse, error) {
	res, err := u.svc.DeleteUserService(p)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserHandler) FindUserByID(ctx context.Context, p *pb.UserID) (*pb.Profile, error) {
	user, err := u.svc.FindUserByIDService(p)
	if err != nil {
		fmt.Println("hiiiiiii")
		return nil, err
	}
	return user, nil
}

func (u *UserHandler) FindUserByEmail(ctx context.Context, p *pb.Email) (*pb.Profile, error) {
	user, err := u.svc.FindUserByEmailService(p)
	if err != nil {
		return nil, err
	}
	return user, nil
}

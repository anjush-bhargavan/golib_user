package handler

import (
	"context"
	"fmt"

	pb "github.com/anjush-bhargavan/golib_user/pkg/pb"
	service_interfaces "github.com/anjush-bhargavan/golib_user/pkg/service/interfaces"
)


type UserHandler struct {
	pb.UserServiceServer
	svc service_interfaces.UserServiceInter
}

func NewUserHandler(svc service_interfaces.UserServiceInter) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) UserSignup(ctx context.Context, p *pb.SignupRequest) (*pb.CommonResponse, error) {
	result, err := u.svc.SignupService(p)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "failed",
			Error:   "user signup error",
			Message: "",
		}, err
	}
	msg := fmt.Sprintf("User created %v", result)
	return &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: msg,
	}, nil
}


func (u *UserHandler) UserLogin(ctx context.Context,p *pb.LoginRequest) (*pb.CommonResponse, error) {
	result, err := u.svc.LoginService(p)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *UserHandler) UserProfile(ctx context.Context,p *pb.UserID) (*pb.Profile, error) {
	result, err := u.svc.ProfileService(p)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *UserHandler) UserProfileUpdate(ctx context.Context, p *pb.ProfileUpdate,) (*pb.Profile, error) {
	result,err := u.svc.UpdateProfileService(p)
	if err != nil {
		return nil,err
	}
	return result,nil
}


package service

import (
	"fmt"

	dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"
	bookpb "github.com/anjush-bhargavan/golib_user/pkg/books/pb"
	pb "github.com/anjush-bhargavan/golib_user/pkg/pb"
	repo_interfaces "github.com/anjush-bhargavan/golib_user/pkg/repo/interfaces"
	service_interfaces "github.com/anjush-bhargavan/golib_user/pkg/service/interfaces"
)

type UserService struct {
	repo   repo_interfaces.UserRepoInter
	client bookpb.BookServicesClient
}

func NewUserService(repo repo_interfaces.UserRepoInter, client bookpb.BookServicesClient) service_interfaces.UserServiceInter {
	return &UserService{
		repo:   repo,
		client: client,
	}
}

func (u *UserService) SignupService(userpb *pb.SignupRequest) (*pb.CommonResponse, error) {
	hashedPassword, err := u.HashPassword(userpb.Password)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "password hashing error",
		}, err
	}
	user := &dom.User{
		FirstName: userpb.Firstname,
		LastName:  userpb.Lastname,
		UserName:  userpb.Firstname + " " + userpb.Lastname,
		DoB:       userpb.Dob,
		Gender:    userpb.Gender,
		Email:     userpb.Email,
		Phone:     userpb.Phone,
		Role:      "user",
		Address:   userpb.Address,
		Password:  hashedPassword,
	}

	err = u.repo.CreateUser(user)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "user signed up error",
		}, err
	}
	return &pb.CommonResponse{
		Status:  "success",
		Error:   "",
		Message: "user signed up successfully",
	}, nil
}

func (u *UserService) LoginService(userpb *pb.LoginRequest) (*pb.CommonResponse, error) {

	user, err := u.repo.FindUserByEmail(userpb.Email)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "error finding user by email",
		}, err
	}

	if err := u.CheckPassword(userpb.Password, user.Password); err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "Invalid password",
		}, err
	}

	if user.Role != "user" {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "invalid role",
		}, err
	}

	token, err := u.GenerateToken(userpb.Email)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "error generating token",
		}, err
	}

	return &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: "Token " + token,
	}, nil
}

func (u *UserService) FindAllUsersService(p *pb.UNoParam) (*pb.Users, error) {
	result, err := u.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	var users []*pb.Profile
	for _, i := range result {
		users = append(users, &pb.Profile{
			Firstname: i.FirstName,
			Lastname:  i.LastName,
			Username:  i.UserName,
			Email:     i.Email,
			Dob:       i.DoB,
			Gender:    i.Gender,
			Phone:     i.Phone,
			Address:   i.Address,
		})
	}
	response := &pb.Users{
		Userlist: users,
	}

	return response, nil
}

func (u *UserService) EditUserService(p *pb.SignupRequest) (*pb.CommonResponse, error) {
	user, err := u.repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}

	user = &dom.User{
		UserID:    uint(user.UserID),
		FirstName: p.Firstname,
		LastName:  p.Lastname,
		UserName:  user.UserName,
		DoB:       p.Dob,
		Gender:    p.Gender,
		Email:     user.Email,
		Phone:     p.Phone,
		Address:   p.Address,
		Password:  user.Password,
	}

	err = u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: "user edited",
	}, nil
}

func (u *UserService) FindUserByEmailService(p *pb.Email) (*pb.Profile, error) {
	user, err := u.repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}
	return &pb.Profile{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Username:  user.UserName,
		Email:     user.Email,
		Dob:       user.DoB,
		Gender:    user.Gender,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (u *UserService) FindUserByIDService(p *pb.UserID) (*pb.Profile, error) {
	user, err := u.repo.FindUserByID(uint(p.Id))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.Profile{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Username:  user.UserName,
		Email:     user.Email,
		Dob:       user.DoB,
		Gender:    user.Gender,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (u *UserService) DeleteUserService(p *pb.UserID) (*pb.CommonResponse, error) {
	err := u.repo.DeleteUser(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Status:  "success",
		Error:   "",
		Message: "user deleted ",
	}, nil
}

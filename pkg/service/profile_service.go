package service

import (
	dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"
	pb "github.com/anjush-bhargavan/golib_user/pkg/pb"
)


func (u *UserService) ProfileService(userpb *pb.UserID) (*pb.Profile,error) {
	user,err := u.repo.FindUserByID(uint(userpb.Id))
	if err != nil {
		return nil,err
	}
	
	response := &pb.Profile{
		Firstname: user.FirstName,
		Lastname: user.LastName,
		Username: user.UserName,
		Dob: user.DoB,
		Gender: user.Gender,
		Email: user.Email,
		Phone: user.Phone,
		Address: user.Address,
	}
	return response,nil
}

func (u *UserService) UpdateProfileService(userpb *pb.ProfileUpdate) (*pb.Profile,error) {
	user,err := u.repo.FindUserByID(uint(userpb.Userid))
	if err != nil {
		return nil,err
	}

	user = &dom.User{
		UserID: uint(userpb.Userid),
		FirstName: userpb.Firstname,
		LastName: userpb.Lastname,
		UserName: userpb.Username,
		DoB: userpb.Dob,
		Gender: userpb.Gender,
		Email: user.Email,
		Phone: userpb.Phone,
		Address: userpb.Address,
		Password: user.Password,
	}
	
	err = u.repo.UpdateUser(user)
	if err != nil {
		return nil,err
	}

	return &pb.Profile{
		Firstname: user.FirstName,
		Lastname: user.LastName,
		Username: user.UserName,
		Email: user.Email,
		Dob: user.DoB,
		Gender: user.Gender,
		Phone: user.Phone,
		Address: user.Address,
	},nil
}
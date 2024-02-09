package service

import (
	"fmt"
	"time"

	dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"
	pb "github.com/anjush-bhargavan/golib_user/pkg/pb"
)

func (u *UserService) FindMembershipByIDService(p *pb.UserID) (*pb.Membership, error) {
	user, err := u.repo.FindMemberByID(uint(p.Id))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.Membership{
		UserId:         uint32(user.UserID),
		SubscriptionId: user.RazorpaySubscriptionID,
		Plan:           user.Plan,
		IsActive:       user.IsActive,
		StartedOn:      user.StartedOn.String(),
		ExpiresAt:      user.ExpiresAt.String(),
	}, nil
}

func (u *UserService) CreateMembershipService(p *pb.Membership) (*pb.CommonResponse, error) {
	// startedOn, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST _", p.StartedOn)
	// if err != nil {
	// 	return nil, err
	// }
	// expiresAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST _", p.ExpiresAt)
	// if err != nil {
	// 	return nil, err
	// }
	userPlan := p.Plan
	var d time.Duration
	if userPlan == "5M" {
		d = time.Hour * 24 * 150
		} else if userPlan == "1Y" {
		d = time.Hour * 24 * 365
	} else if userPlan == "3Y" {
		d = time.Hour * 24 * 1095
	}

	membership := &dom.Membership{
		UserID:                 uint(p.UserId),
		RazorpaySubscriptionID: p.SubscriptionId,
		Plan:                   p.Plan,
		IsActive:               true,
		StartedOn:              time.Now(),
		ExpiresAt:              time.Now().Add(d),
	}

	err := u.repo.CreateMembership(membership)
	if err != nil {
		return nil, err
	}

	return &pb.CommonResponse{
		Status: "Success",
		Error: "",
		Message: "Membership created",
	},nil
}

func (u *UserService) UpdateMembershipService(p *pb.Membership) (*pb.CommonResponse,error) {
	startedOn, err := time.Parse(time.RFC3339, p.StartedOn)
	if err != nil {
		return nil, err
	}
	expiresAt, err := time.Parse(time.RFC3339, p.ExpiresAt)
	if err != nil {
		return nil, err
	}

	membership := &dom.Membership{
		UserID:                 uint(p.UserId),
		RazorpaySubscriptionID: p.SubscriptionId,
		Plan:                   p.Plan,
		IsActive:               true,
		StartedOn:              startedOn,
		ExpiresAt:              expiresAt,
	}

	err = u.repo.UpdateMember(membership)
	if err != nil {
		return nil, err
	}

	return &pb.CommonResponse{
		Status: "Success",
		Error: "",
		Message: "Membership updated",
	},nil
}

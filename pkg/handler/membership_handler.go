package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/golib_user/pkg/pb"
)

func (u *UserHandler) CreateMembership(ctx context.Context, p *pb.Membership) (*pb.CommonResponse, error) {
	result, err := u.svc.CreateMembershipService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) FindMembershipByID(ctx context.Context, p *pb.UserID) (*pb.Membership, error) {
	result, err := u.svc.FindMembershipByIDService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) UpdateMembeship(ctx context.Context, p *pb.Membership) (*pb.CommonResponse, error) {
	result, err := u.svc.UpdateMembershipService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}
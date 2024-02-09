package repo_interfaces

import dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"

// UserRepo defines the methods in UserRepo
type UserRepoInter interface {
	CreateUser(user *dom.User) error
	FindUserByID(id uint) (*dom.User, error)
	UpdateUser(user *dom.User) error
	FindUserByEmail(email string) (*dom.User, error)
	DeleteUser(id uint) error
	FindAllUsers() ([]*dom.User, error)

	CheckMembership(userID uint) bool
	CreateMembership(user *dom.Membership) error
	FindMemberByID(id uint) (*dom.Membership, error)
	UpdateMember(user *dom.Membership) error
}

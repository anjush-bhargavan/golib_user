package repo

import (
	dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"
	"gorm.io/gorm"
)

// UserRepository for connecting db to UserRepoInter methods
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepo creates an instance of user repo
func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a newuser in database else returns error
func (u *UserRepository) CreateUser(user *dom.User) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindUserByID finds the user in the database using ID
func (u *UserRepository) FindUserByID(id uint) (*dom.User, error) {
	var user *dom.User
	err := u.DB.First(&user, id).Error
	return user, err
}

// UpdateUser updates user details in the databse
func (u *UserRepository) UpdateUser(user *dom.User) error {
	err := u.DB.Save(&user).Error
	return err
}

// DeleteUser deletes a user using the id
func (u *UserRepository) DeleteUser(id uint) error {
	err := u.DB.Delete(&dom.User{}, id).Error
	return err
}

// FindUserByEmail finds the user by email
func (u *UserRepository) FindUserByEmail(email string) (*dom.User, error) {
	var user *dom.User
	err := u.DB.Where("email = ?",email).First(&user,).Error
	return user, err
}

//FindAllUsers finds the all the user details in the database
func (u *UserRepository) FindAllUsers() ([]*dom.User, error) {
	var users []*dom.User
	err := u.DB.Find(&users).Error
	return users, err
}

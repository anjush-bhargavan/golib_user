package repo

import dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"

// FindUserByID finds the user in the database using ID
func (u *UserRepository) CheckMembership(id uint) bool {
	var membership *dom.Membership
	err := u.DB.First(&membership, id).Error
	return err == nil
}

// CreateMembership creates a member in database else returns error
func (u *UserRepository) CreateMembership(user *dom.Membership) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindMemberByID finds the member in the database using ID
func (u *UserRepository) FindMemberByID(id uint) (*dom.Membership, error) {
	var user *dom.Membership
	err := u.DB.Where("user_id = ?",id).First(&user).Error
	return user, err
}

// UpdateMember updates membership details in the database
func (u *UserRepository) UpdateMember(user *dom.Membership) error {
	err := u.DB.Where("user_id = ?",user.UserID).Save(&user).Error
	return err
}
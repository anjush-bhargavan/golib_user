package service

import "golang.org/x/crypto/bcrypt"

// HashPassword will hash the password of user
func (user *UserService) HashPassword(password string) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "",err
	}
	Password := string(bytes)
	return Password,nil
}


//CheckPassword function will check the provided password with users password
func (user *UserService) CheckPassword(providedPassword,password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

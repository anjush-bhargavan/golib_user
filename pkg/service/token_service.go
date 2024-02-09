package service

import (
	"time"

	dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"
	"github.com/golang-jwt/jwt"
)

//GenerateToken to generate jwt token
func (u *UserService)GenerateToken(userEmail string) (string,error){
	user,err := u.repo.FindUserByEmail(userEmail)
	if err != nil {
		return "",err
	}

	isActive := u.repo.CheckMembership(user.UserID)

	claims:= &dom.Claims{
		UserID: user.UserID,
		Email: user.Email,
		Role: user.Role,
		IsActive: isActive,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*24*10).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}


	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	secretKey := []byte("101101")

	tokenString,err :=token.SignedString(secretKey)
	if err != nil{
		return "",err
	}

	return tokenString,nil
}



package dom

import "github.com/golang-jwt/jwt"

// Claims that passing via jwt token
type Claims struct {
	UserID   uint `json:"userid"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool
	Claims   jwt.StandardClaims
}

// Valid implements jwt.Claims.
func (*Claims) Valid() error {
	panic("unimplemented")
}

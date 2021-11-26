package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// type JwtWrapper struct {
// 	SecretKey string
// 	ExpirationHours int64
// }
var mySecretKey = []byte("verysecretkey")

type jwtClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(username, role string) *jwtClaim {
	claims := &jwtClaim{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 5).Unix(),
		},
	}
	return claims
}

func (jcl *jwtClaim) Create() (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jcl)

	signedToken, err = token.SignedString(mySecretKey)

	if err != nil {
		return
	}
	return
}

func ValidateToken(signedToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySecretKey), nil
		})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*jwtClaim)

	return token.Valid, claims.Role

}

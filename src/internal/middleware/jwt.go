package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/eirene/src/config"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
)

type jwtclaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user entity.UserLogin) (string, error) {
	conf := config.GetConfig()
	time := jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	claims := &jwtclaims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(conf.Server.Key)
	if err != nil {
		return "", helper.InternalServerError("cannot generate token")
	}

	return tokenString, nil
}

func CheckToken(token string) (*jwtclaims, error) {
	conf := config.GetConfig()
	tokens, _ := jwt.ParseWithClaims(token, &jwtclaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Server.Key), helper.Forbidden("this is strict page")
	})

	claim, err := tokens.Claims.(*jwtclaims)
	if !err {
		return nil, helper.Unauthorize("sorry you not have token")
	}

	return claim, nil

}

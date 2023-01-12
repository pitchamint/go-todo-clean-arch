package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

func CreateJWTToken(user models.LoginUser) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 6).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Email
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = exp
	jwtTokenstr, err := token.SignedString([]byte(jwtkey))
	if err != nil {
		return "", 0, err
	}
	return jwtTokenstr, exp, nil
}

// get
const (
	jwtkey = "secret"
)

func ValidateToken(jwtTokenStr string, jwtkey string) (status bool, err error) {
	token, err := extractToken(jwtTokenStr, jwtkey)
	if err != nil {
		return false, err
	}
	if !token.Valid {
		err = fmt.Errorf("invalid Jwt-token")
		return false, err
	}
	return true, err
}

func ValidateAndExtractEmail(tokenString string, jwtkey string) (email string, err error) {
	token, err := extractToken(tokenString, jwtkey)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		email, ok := claims["name"].(string)
		if !ok {
			return "", fmt.Errorf("not access email")
		}
		return email, nil
	}
	return "", fmt.Errorf("error when parsing claims")
}

func extractToken(tokenString, jwtAcckey string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method not matched")
		}
		return []byte(jwtAcckey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

package jwt_service

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id int) string {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	jwtExpireDay, _ := strconv.Atoi(os.Getenv("JWT_DAY_EXPIRE_TIME"))
	token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(jwtExpireDay)).Unix(),
	}
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString
}

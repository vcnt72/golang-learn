package service

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name string , admin bool) string
	ValidateToken(name string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer string
}


func getSecretKey() string {
	secret := "secret"
	return secret
}


func NewJwtService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer: "vincent.sanjaya.com",
	}
}

func (jwtSrv *jwtService) GenerateToken(username string , admin bool) string {
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtSrv.secretKey))

	if err != nil {
		panic(err)
	}
	return t
}


func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
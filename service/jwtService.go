package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserdId string `json:"userId"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JWTService {
	return &jwtService{
		issuer:    "odhiahmad",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")

	if secretKey != "" {
		secretKey = "odhiahmad"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserId string) string {
	claims := &jwtCustomClaim{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

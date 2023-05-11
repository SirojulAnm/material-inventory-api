package auth

import (
	"errors"
	"strconv"
	"time"

	"log"

	"tripatra-api/db"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(tokenStr string) (*CustomClaims, error)
	DeleteToken(userID int) (string, error)
}

type jwtService struct {
	SecretKey string
}

type CustomClaims struct {
	jwt.StandardClaims
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	now := time.Now().UTC()
	expirationTime := time.Now().Add(time.Hour * 24)
	//24 * time.Hour * 30
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  now.Unix(),
			Subject:   strconv.Itoa(userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		log.Printf("Error while encode jwt %s", err.Error())
		return "", err
	}

	redis, err := db.ConnRedis()
	if err != nil {
		return "", err
	}

	// save token in redis
	err = redis.Set(strconv.Itoa(userID), tokenStr, expirationTime.Sub(now)).Err()
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *jwtService) ValidateToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		log.Printf("Error while parse jwt %s", err.Error())
		return nil, errors.New("Unauthorized, Error while parse jwt")
	}

	claims, ok := token.Claims.(*CustomClaims)

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, err
	}

	redis, err := db.ConnRedis()
	if err != nil {
		return nil, err
	}

	// check if token is in redis
	storedToken, err := redis.Get(strconv.Itoa(userID)).Result()
	if err != nil {
		return nil, err
	}

	// compare stored token with incoming token
	if tokenStr != storedToken {
		return nil, errors.New("Token doesn't match data in redis")
	}

	if ok && token.Valid {
		return claims, nil
	} else {
		log.Println("Error while claims jwt", err.Error())
		return nil, errors.New("Unauthorized, Error while claims jwt")
	}
}

func (s *jwtService) DeleteToken(userID int) (string, error) {
	redis, err := db.ConnRedis()
	if err != nil {
		return "", err
	}
	storedToken, err := redis.Get(strconv.Itoa(userID)).Result()

	// Menghapus token dari Redis
	err = redis.Del(strconv.Itoa(userID)).Err()
	if err != nil {
		return "", err
	}

	return storedToken, nil
}

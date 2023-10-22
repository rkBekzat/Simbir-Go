package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"vtb_api/internal/entities"
	"vtb_api/internal/repository"
)

const (
	salt       = "desgerfr3241rfgfqwer"
	signingKey = "dqwfdqwdewvafsdvfcs12erf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type auth struct {
	repo repository.Authorization
}

func NewAuthorization(repo repository.Authorization) Auth {
	return &auth{repo: repo}
}

func (a *auth) CreateUser(user *entities.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *auth) GenerateToken(username, password string) (string, error) {
	user, err := a.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (a *auth) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}

func (a *auth) Update(id int, username, password string) error {
	user, _ := a.repo.GetUserById(id)
	if user.Username == username && password == "" {
		return errors.New("Nothing updated")
	}
	_, err := a.repo.GetUserByUsername(username)
	if err != nil {
		password = generatePasswordHash(password)
		return a.repo.UpdateUser(id, username, password)
	}
	return errors.New("This username chosen by other user")
}

func (a *auth) Information(id int) (*entities.User, error) {
	return a.repo.GetUserById(id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

package service

import (
	"crypto/sha1"
	"fmt"
	"vtb_api/internal/entities"
	"vtb_api/internal/repository"
)

const salt = "desgerfr3241rfgfqwer"

type auth struct {
	repo repository.Authorization
}

func NewAuthorization(repo repository.Authorization) Auth {
	return &auth{repo: repo}
}

func (a *auth) CreateUser(user entities.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

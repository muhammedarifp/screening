package usecases

import (
	"errors"
	"regexp"
	"screening/repo"
)

type UseCase struct {
	Repo repo.DataBaseInterface
}

var (
	ErrInvalidEmail = errors.New("invalid email")
)

func New(repo repo.DataBaseInterface) *UseCase {
	return &UseCase{
		Repo: repo,
	}
}

func (u *UseCase) CreateUser(name, email string) (int64, error) {
	if !isValidEmail(email) {
		return 0, ErrInvalidEmail
	}
	return u.Repo.CreateNewUser(name, email)
}

func (u *UseCase) UpdateUser(id int64, name, email string) (int64, error) {
	if !isValidEmail(email) {
		return 0, ErrInvalidEmail
	}
	return u.Repo.UpdateUser(id, name, email)
}

// is email valid
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

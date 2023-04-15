package repository

import (
	"errors"

	"github.com/krisnasw/go-grst/examples/users-api/entity"
)

var ErrUserNotFound = errors.New("User not found")
var ErrUserAlreadyExist = errors.New("User already exist")

type UserRepository interface {
	GetById(id int) (*entity.User, error)
	GetProfileById(id int) (*entity.UserProfile, error)
}

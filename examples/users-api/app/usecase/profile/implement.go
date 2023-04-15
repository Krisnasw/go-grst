package profile

import (
	"errors"
	"fmt"

	"github.com/krisnasw/go-grst/examples/users-api/app/repository"
	"github.com/krisnasw/go-grst/examples/users-api/entity"
)

type usecase struct {
	user_repo repository.UserRepository
}

func New(user_repo repository.UserRepository) UseCase {
	return &usecase{user_repo: user_repo}
}

func (uc *usecase) GetProfile(id int) (*entity.UserProfile, error) {
	data, err := uc.user_repo.GetProfileById(id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}

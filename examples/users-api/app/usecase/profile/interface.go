package profile

import "github.com/krisnasw/go-grst/examples/users-api/entity"

type UseCase interface {
	GetProfile(id int) (*entity.UserProfile, error)
}

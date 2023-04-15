package profile

import "github.com/krisnasw/cdd/examples/users-api/entity"

type UseCase interface {
	GetProfile(id int) (*entity.UserProfile, error)
}

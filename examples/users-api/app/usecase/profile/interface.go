package profile

import "github.com/herryg91/cdd/examples/users-api/entity"

type UseCase interface {
	GetProfile(id int) (*entity.UserProfile, error)
}

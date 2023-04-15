package search_usecase

import "github.com/herryg91/cdd/examples/location-api/entity"

type UseCase interface {
	Search(keyword string) ([]entity.CityProfile, error)
}

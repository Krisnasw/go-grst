package search_usecase

import "github.com/krisnasw/go-grst/examples/location-api/entity"

type UseCase interface {
	Search(keyword string) ([]entity.CityProfile, error)
}

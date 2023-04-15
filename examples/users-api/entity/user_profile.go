package entity

import (
	pbUsers "github.com/krisnasw/go-grst/examples/users-api/handler/grst/users"
)

func (UserProfile) FromUser(u User, provinceName string) UserProfile {
	return UserProfile{
		Id:           u.Id,
		Name:         u.Name,
		ProvinceId:   u.ProvinceId,
		ProvinceName: provinceName,
	}
}
func (u *UserProfile) ToPbUserProfile() *pbUsers.UserProfile {
	return &pbUsers.UserProfile{
		Id:           int32(u.Id),
		Name:         u.Name,
		ProvinceId:   int32(u.ProvinceId),
		ProvinceName: u.ProvinceName,
	}
}

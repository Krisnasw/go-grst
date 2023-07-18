package entity

func (CityProfile) FromCity(in City, provinceName string) CityProfile {
	return CityProfile{
		Id:           in.Id,
		Name:         in.Name,
		ProvinceId:   in.ProvinceId,
		ProvinceName: provinceName,
	}
}

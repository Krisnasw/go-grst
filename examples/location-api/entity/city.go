package entity

type Cities []*City

func (c Cities) GetProvinceIds() []int {
	provinceIds := []int{}
	for _, data := range c {
		provinceIds = append(provinceIds, data.ProvinceId)
	}
	return provinceIds
}

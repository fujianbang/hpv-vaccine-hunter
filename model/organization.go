package model

// VaccineItem 疫苗
type VaccineItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ImgUrl      string `json:"imgUrl"`
	VaccineCode string `json:"vaccineCode"` // 疫苗编号
	VaccineName string `json:"vaccineName"` // 疫苗名称
	Address     string `json:"address"`     // 机构地址
	StartTime   string `json:"startTime"`   // 开始时间
	Stock       int    `json:"stock"`       // 库存
}

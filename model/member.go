package model

type Member struct {
	Id           int    `json:"id"`           // 唯一ID
	UserId       int    `json:"userId"`       // 用户ID
	Name         string `json:"name"`         // 姓名
	IdCardNo     string `json:"idCardNo"`     // 身份证号
	Birthday     string `json:"birthday"`     // 生日
	Sex          int    `json:"sex"`          // 性别 2:女
	RegionCode   string `json:"regionCode"`   // 地区码 例如: 510109
	Address      string `json:"address"`      // 地址
	IsDefault    int    `json:"isDefault"`    // 是否默认
	RelationType int    `json:"relationType"` // 与本人关系 1:本人
	CreateTime   string `json:"createTime"`   // 创建时间
	ModifyTime   string `json:"modifyTime"`   // 修改时间
	Yn           int    `json:"yn"`           // Unknown
	IdCardType   int    `json:"idCardType"`   // 身份证类型 1:身份证
}

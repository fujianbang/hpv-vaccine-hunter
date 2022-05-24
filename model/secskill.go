package model

type CheckStockResult struct {
	Stock           int   `json:"stock"` //	状态
	ServerTimestamp int64 `json:"st"`    // 服务器时间戳 毫秒
}

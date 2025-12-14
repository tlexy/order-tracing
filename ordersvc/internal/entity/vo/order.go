package vo

// OrderCreateReq 创建订单请求参数
type OrderCreateReq struct {
	UserID    int64   `json:"user_id"`    // 用户ID
	ProductID int64   `json:"product_id"` // 产品ID
	Amount    float64 `json:"amount"`     // 订单金额
}

// OrderCreateResp 创建订单响应参数
type OrderCreateResp struct {
	OrderID string `json:"order_id"` // 订单ID
	Status  string `json:"status"`   // 订单状态
}

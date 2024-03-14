package TransactionDbCtrl

import (
	"C2CTranstion_Server/OrderDbCtrl"
	"time"
)

type CreateOrderInsert struct {
	PurchasersID uint32  `json:"purchasers_id"`
	SellerID     uint32  `json:"seller_id"`
	SuretyID     uint32  `json:"surety_id"`
	OrderStatus  uint8   `json:"order_status"`
	ProductType  uint8   `json:"product_type"`
	Prices       float64 `json:"prices"`
	Count        int     `json:"count"`
	TotalPrice   float64 `json:"total_price"`
	IsRetail     bool    `json:"is_retail"`
}

func CreateOrder(request CreateOrderInsert) (err error) {
	timeNow := time.Now()
	errS := Db.Create(&OrderDbCtrl.Order{
		PurchasersID: request.PurchasersID,
		SellerID:     request.SellerID,
		SuretyID:     request.SuretyID,
		OrderStatus:  request.OrderStatus,
		ProductType:  request.ProductType,
		Prices:       request.Prices,
		Count:        request.Count,
		TotalPrice:   request.TotalPrice,
		IsRetail:     request.IsRetail,
		Time:         timeNow,
	}).Error

	return errS
}

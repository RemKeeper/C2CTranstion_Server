package OrderDbCtrl

import "time"

type Order struct {
	OrderID       uint32    `gorm:"column:order_id; primaryKey; autoIncrement" json:"order_id"`
	PurchasersID  uint32    `gorm:"column:purchasers_id" json:"purchasers_id"`
	SellerID      uint32    `gorm:"column:seller_id" json:"seller_id"`
	SuretyID      uint32    `gorm:"column:surety_id" json:"surety_id"`
	OrderStatus   uint8     `gorm:"column:order_status" json:"order_status"`
	FreezeReasons string    `gorm:"column:freeze_reasons" json:"freeze_reasons"`
	ProductType   uint8     `gorm:"column:product_type" json:"product_type"`
	Prices        float64   `gorm:"column:prices" json:"prices"`
	Count         int       `gorm:"column:count" json:"count"`
	IsDispatched  bool      `gorm:"column:is_dispatched" json:"is_dispatched"`
	TotalPrice    float64   `gorm:"column:total_price" json:"total_price"`
	IsRetail      bool      `gorm:"column:is_retail" json:"is_retail"`
	Time          time.Time `gorm:"column:time" json:"time"`
}

func (Order) TableName() string {
	return "order"
}

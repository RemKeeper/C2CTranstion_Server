package OrderDbCtrl

import (
	"time"
)

type SubOrder struct {
	PrentOrderID   int       `gorm:"primary_key;column:prent_order_id" json:"prent_order_id"`
	SubOrderID     int       `gorm:"primary_key;column:sub_order_id" json:"sub_order_id"`
	Count          int       `gorm:"column:count" json:"count"`
	PurchasersID   int       `gorm:"column:purchasers_id" json:"purchasers_id"`
	SellerID       int       `gorm:"column:seller_id" json:"seller_id"`
	ProductType    int       `gorm:"column:product_type" json:"product_type"`
	TotalPrice     float64   `gorm:"column:total_price" json:"total_price"`
	SubOrderStatus int       `gorm:"column:sub_order_status" json:"sub_order_status"`
	Time           time.Time `gorm:"column:time" json:"time"`
}

func (SubOrder) TableName() string {
	return "sub_order"
}

func CreateSubOrder(subOrder SubOrder) error {
	return Db.Create(&subOrder).Error
}

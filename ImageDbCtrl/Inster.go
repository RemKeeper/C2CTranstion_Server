package ImageDbCtrl

import (
	"C2CTranstion_Server/DbInit"
	"errors"
	"time"
)

type OrderImg struct {
	OrderID       int       `gorm:"primary_key;comment:'订单ID'"`
	PurchasersImg string    `gorm:"type:varchar(255);comment:'买家截图'"`
	SellerImg     string    `gorm:"type:varchar(255);comment:'卖家截图'"`
	Time          time.Time `gorm:"comment:'创建时间'"`
}

var Db = DbInit.Db

func (OrderImg) TableName() string {
	return "order_img"
}

func CreateOrderImg(orderImg OrderImg) error {
	var orderImg1 OrderImg
	Db.Where("order_id = ?", orderImg.OrderID).First(&orderImg1)
	switch {
	case orderImg1 == OrderImg{}:
		err := Db.Create(orderImg).Error
		if err != nil {
			return err
		}
	case orderImg1.PurchasersImg == "":
		err := Db.Model(&orderImg1).Update("purchasers_img", orderImg.PurchasersImg).Error
		if err != nil {
			return err
		}
	case orderImg1.SellerImg == "":
		err := Db.Model(&orderImg1).Update("seller_img", orderImg.SellerImg).Error
		if err != nil {
			return err
		}
	default:
		return errors.New("数据已存在")
	}
	return nil
}

func GetOrderImgByOrderID(orderID int) (*OrderImg, error) {
	var orderImg OrderImg
	err := Db.Where("order_id = ?", orderID).First(&orderImg).Error
	if err != nil {
		return nil, err
	}
	return &orderImg, nil
}

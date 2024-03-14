package TransactionDbCtrl

import (
	"C2CTranstion_Server/OrderDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"time"
)

func init() {
	ticker := time.NewTicker(2 * time.Minute)
	ticker2 := time.NewTicker(7 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				CancelTimeOutOrder()
			case <-ticker2.C:
				_ = CancelTimeOutShelfOrder()
			}
		}
	}()
}

func CancelTimeOutOrder() {
	//查询并删除与当前时间相差两分钟 且状态为1的订单
	var order []OrderDbCtrl.Order
	Db.Where("time < ?", time.Now().Add(-2*time.Minute)).Where("order_status = ?", 1).Find(&order).Delete(&OrderDbCtrl.Order{})
}

func CancelTimeOutShelfOrder() error {
	//查询并删除与当前时间相差七分钟 且状态为0的订单,且未创建子订单的  订单
	var order []OrderDbCtrl.Order
	err := Db.Joins("LEFT JOIN sub_order ON order.order_id = sub_order.prent_order_id").
		Where("time < ?", time.Now().Add(-7*time.Minute)).
		Where("order_status = ?", 0).
		Find(&order).Delete(&OrderDbCtrl.Order{})
	if err != nil {
		return err.Error
	}

	for _, v := range order {
		switch v.ProductType {
		case 0:
			UserDbCtrl.UserTransactionEditCoinBalance(true, uint(v.SellerID), v.Count)
		case 1:
			UserDbCtrl.UserTransactionEditCoinBalance(true, uint(v.PurchasersID), int(v.TotalPrice))

		}
	}
	return nil
}

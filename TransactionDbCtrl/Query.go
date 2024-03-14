package TransactionDbCtrl

import (
	"C2CTranstion_Server/OrderCtrl"
	"C2CTranstion_Server/OrderDbCtrl"
)

func GetNormalOrder() ([]OrderCtrl.OrderResponse, error) {
	var order []OrderDbCtrl.Order
	Db.Where("order_status = ?", 0).Find(&order)

	var orderResponse []OrderCtrl.OrderResponse
	for _, v := range order {
		orderResponse = append(orderResponse, OrderCtrl.OrderResponse{
			OrderId:       v.OrderID,
			PurchasersId:  v.PurchasersID,
			SellerId:      v.SellerID,
			SuretyId:      v.SuretyID,
			OrderStatus:   v.OrderStatus,
			FreezeReasons: v.FreezeReasons,
			ProductType:   v.ProductType,
			Prices:        v.Prices,
			Count:         v.Count,
			TotalPrice:    v.TotalPrice,
			Time:          v.Time,
		})
	}
	return orderResponse, nil
}

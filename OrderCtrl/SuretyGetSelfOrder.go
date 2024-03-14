package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
)

func SuretyGetSelfOrder(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "请先登录",
		})
		return
	}
	var OrderResponseList []OrderResponse
	OrderList, err := OrderDbCtrl.SuretyGetSelfOrder(cookie)
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{
			Code: 500,
			Msg:  "获取订单失败",
		})
		return
	}

	for _, Order := range OrderList {
		var OrderResponse OrderResponse
		OrderResponse.OrderId = Order.OrderID
		OrderResponse.PurchasersId = Order.PurchasersID
		OrderResponse.SellerId = Order.SellerID
		OrderResponse.SuretyId = Order.SuretyID
		OrderResponse.OrderStatus = Order.OrderStatus
		OrderResponse.FreezeReasons = Order.FreezeReasons
		OrderResponse.ProductType = Order.ProductType
		OrderResponse.Prices = Order.Prices
		OrderResponse.Count = Order.Count
		OrderResponse.IsDispatched = Order.IsDispatched
		OrderResponse.TotalPrice = Order.TotalPrice
		OrderResponse.Time = Order.Time
		OrderResponseList = append(OrderResponseList, OrderResponse)
	}
	c.JSON(200, OrderResponseList)
}

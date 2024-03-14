package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
)

func GetFreezeOrder(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return
	}
	var OrderResponseList []OrderResponse
	OrderList, err := OrderDbCtrl.UserGetFreezeOrder(cookie)
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{
			Code: 500,
			Msg:  "获取订单列表失败 ",
		})
		return
	} else {
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
			OrderResponse.TotalPrice = Order.TotalPrice
			OrderResponse.Time = Order.Time
			OrderResponseList = append(OrderResponseList, OrderResponse)
		}
		c.JSON(200, OrderResponseList)
	}

}

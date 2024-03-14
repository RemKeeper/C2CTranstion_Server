package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderCtrl"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetOrderById(c *gin.Context) {
	OrderIdStr, ok := c.GetQuery("order_id")
	if !ok {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "无法获取订单id",
		})
		return
	}
	OrderId, err := strconv.Atoi(OrderIdStr)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "无法获取订单id",
		})
		return
	}
	order, err := OrderDbCtrl.GetOrderById(uint32(OrderId))
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "无法获取订单信息",
		})
		return
	}
	OrderResponse := OrderCtrl.OrderResponse{
		OrderId:       order.OrderID,
		PurchasersId:  order.PurchasersID,
		SellerId:      order.SellerID,
		SuretyId:      order.SuretyID,
		OrderStatus:   order.OrderStatus,
		FreezeReasons: order.FreezeReasons,
		ProductType:   order.ProductType,
		Prices:        order.Prices,
		Count:         order.Count,
		IsDispatched:  order.IsDispatched,
		TotalPrice:    order.TotalPrice,
		Time:          order.Time,
	}

	c.JSON(200, OrderResponse)

}

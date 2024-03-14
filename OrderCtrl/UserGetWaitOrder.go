package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
	"time"
)

type OrderResponse struct {
	OrderId       uint32    `json:"order_id"`
	PurchasersId  uint32    `json:"purchasers_id"`
	SellerId      uint32    `json:"seller_id"`
	SuretyId      uint32    `json:"surety_id"`
	OrderStatus   uint8     `json:"order_status"`
	FreezeReasons string    `json:"freeze_reasons"`
	ProductType   uint8     `json:"product_type"`
	Prices        float64   `json:"prices"`
	Count         int       `json:"count"`
	IsDispatched  bool      `gorm:"column:Is_Dispatched" json:"is_dispatched"`
	TotalPrice    float64   `json:"total_price"`
	Time          time.Time `json:"time"`
}

func GetWaitOrder(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return
	}
	var OrderResponseList []OrderResponse = make([]OrderResponse, 0)
	OrderList, err := OrderDbCtrl.UserGetWaitOrder(cookie)
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
			OrderResponse.IsDispatched = Order.IsDispatched
			OrderResponse.TotalPrice = Order.TotalPrice
			OrderResponse.Time = Order.Time
			OrderResponseList = append(OrderResponseList, OrderResponse)
		}
		c.JSON(200, OrderResponseList)
	}
}

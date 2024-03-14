package DividendPercentageICtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/DividendPercentageDbCtrl"
	"C2CTranstion_Server/OrderDbCtrl"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetOrderNeedPay(c *gin.Context) {
	OrderIdStr, exit := c.GetQuery("order_id")
	if !exit {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "参数错误",
		})
		return
	}
	OrderId, err := strconv.Atoi(OrderIdStr)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "参数错误",
		})
		return
	}
	order, err := OrderDbCtrl.GetOrderById(uint32(OrderId))
	if err != nil {
		return
	}

	c.JSON(200, CommunicationStructure.Message{
		Code: 200,
		Msg:  fmt.Sprint(DividendPercentageDbCtrl.OrderNeedPay(order)),
	})
}

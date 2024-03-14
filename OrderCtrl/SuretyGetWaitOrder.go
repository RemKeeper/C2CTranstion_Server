package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
)

func SuretyGetWaitOrder(c *gin.Context) {
	order, err := OrderDbCtrl.SuretyGetWaitOrder()
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "查询失败 " + err.Error(),
		})
		return
	}
	c.JSON(200, order)
}

package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOrder(c *gin.Context) {
	sortStr := c.Query("sort")
	OrderList, err := OrderDbCtrl.GetAllOrder(sortStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "获取订单列表失败 " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, OrderList)
	}
}

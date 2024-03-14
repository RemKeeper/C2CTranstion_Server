package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserGetAllOrder(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return
	}
	OrderList, err := OrderDbCtrl.UserGetAllOrder(cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "获取订单列表失败 " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, OrderList)
	}
}

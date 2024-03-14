package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

type OrderRequest struct {
	OrderID int `json:"order_id"`
}

func DeleteOrder(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败 " + err.Error(),
		})
		return
	}
	var deleteOrderRequest OrderRequest
	err = json.Unmarshal(data, &deleteOrderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败 " + err.Error(),
		})
		return
	}
	err = OrderDbCtrl.DeleteOrder(deleteOrderRequest.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "删除订单失败 " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "删除订单成功",
	})
}

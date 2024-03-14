package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func FreezeOrder(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败 " + err.Error(),
		})
		return
	}
	var freezeOrderRequest OrderRequest
	err = json.Unmarshal(data, &freezeOrderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败 " + err.Error(),
		})
		return
	}
	err = OrderDbCtrl.FreezeOrder(freezeOrderRequest.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "冻结订单失败 " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, CommunicationStructure.Message{
			Code: http.StatusOK,
			Msg:  "冻结订单成功",
		})
	}

}

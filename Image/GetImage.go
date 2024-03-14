package Image

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/ImageDbCtrl"
	"C2CTranstion_Server/OrderCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func GetOrderImage(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败 " + err.Error(),
		})
		return
	}
	var getOrderImageRequest OrderCtrl.OrderRequest
	err = json.Unmarshal(data, &getOrderImageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败 " + err.Error(),
		})
		return
	}
	OrderImage, err := ImageDbCtrl.GetOrderImgByOrderID(getOrderImageRequest.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败 " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &OrderImage)

}

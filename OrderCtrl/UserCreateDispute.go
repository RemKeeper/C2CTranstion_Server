package OrderCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func CreateDispute(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败 " + err.Error(),
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败 ",
		})
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败 " + err.Error(),
		})
		return
	}
	var CreateDisputeRequest OrderRequest
	err = json.Unmarshal(data, &CreateDisputeRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败 " + err.Error(),
		})
		return
	}

	err = OrderDbCtrl.UserFreezeOrder(CreateDisputeRequest.OrderID, int(user.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "冻结订单失败 ",
		})
		return
	}
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "冻结订单成功",
	})

}

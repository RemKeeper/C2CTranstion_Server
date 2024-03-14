package transaction

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/TransactionDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func ConfirmSurety(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		return
	}
	var CarrySuretyRaw CarryRequest
	err = json.Unmarshal(data, &CarrySuretyRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败",
		})
		return
	}
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取cookie失败",
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "获取用户信息失败",
		})
		return
	}
	TransactionDbCtrl.ConfirmSurety(CarrySuretyRaw.OrderID, user.UserID)
}

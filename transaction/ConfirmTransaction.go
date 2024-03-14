package transaction

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/TransactionDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func ConfirmTransaction(c *gin.Context) {
	var request CarryTransactionRequest
	data, err := c.GetRawData()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &request)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析json失败",
		})
		return
	}
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取cookie失败",
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取用户信息失败",
		})
		return
	}
	TransactionDbCtrl.ConfirmTransaction(request.OrderID, user.UserID)
	c.JSON(200, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "确认成功",
	})
}

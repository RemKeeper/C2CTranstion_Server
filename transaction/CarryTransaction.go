package transaction

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"C2CTranstion_Server/TransactionDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

type CarryTransactionRequest struct {
	OrderID uint32 `json:"order_id"`
}

func CarryTransaction(c *gin.Context) {
	var request CarryTransactionRequest
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "未获取到有效载荷",
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
	err = json.Unmarshal(data, &request)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析json失败",
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

	if user.PriLabel >= 3 {

	}

	order, err := OrderDbCtrl.GetOrderById(request.OrderID)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  fmt.Sprint("获取订单信息失败", request.OrderID),
		})
		return
	}
	if order.OrderStatus != 0 {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "非法操作,该订单已被接取",
		})
		return
	}

	if uint(order.SellerID) == user.UserID || uint(order.PurchasersID) == user.UserID {

		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "非法操作,您不能接取自己的订单",
		})
		return
	}

	err = TransactionDbCtrl.CarryTransaction(order, user)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  fmt.Sprint("交易失败", err),
		})
		return
	}
	c.JSON(200, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  fmt.Sprint("交易状态更新成功,订单号:", request.OrderID, " 订单总金额:", order.TotalPrice, "\n请五分钟内完成交易\n如果您拥有余额，此单将自动余额交易"),
	})

}

package transaction

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/OrderDbCtrl"
	"C2CTranstion_Server/TransactionDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"C2CTranstion_Server/Utils"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"math"
	"net/http"
)

type CarryRequest struct {
	OrderID uint32 `json:"order_id"`
}

func CarrySurety(c *gin.Context) {
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
	order, err := OrderDbCtrl.GetOrderById(CarrySuretyRaw.OrderID)
	if err != nil {
		return
	}
	if uint(order.PurchasersID) == user.UserID || uint(order.SellerID) == user.UserID {
		c.JSON(http.StatusForbidden, CommunicationStructure.Message{
			Code: http.StatusForbidden,
			Msg:  "您不能担保有关于自己的订单",
		})
		return
	}

	switch order.ProductType {
	case 0:
		if user.CoinBalance < order.Count {
			if user.PriLabel < 4 {
				c.JSON(http.StatusForbidden, CommunicationStructure.Message{
					Code: http.StatusForbidden,
					Msg:  "您没有足够的担保金",
				})
				return
			} else {
				if Utils.Abs(user.CoinBalance-order.Count) >= user.CoinBalanceQuota {
					c.JSON(http.StatusForbidden, CommunicationStructure.Message{
						Code: http.StatusForbidden,
						Msg:  "您没有足够的担保金(可负额度不足)",
					})
					return
				}
			}
		}
	case 1:
		if user.Balance < order.TotalPrice {
			if user.PriLabel < 4 {
				c.JSON(http.StatusForbidden, CommunicationStructure.Message{
					Code: http.StatusForbidden,
					Msg:  "您没有足够的担保金",
				})
				return
			} else {
				if math.Abs(user.Balance-order.TotalPrice) >= user.BalanceQuota {
					c.JSON(http.StatusForbidden, CommunicationStructure.Message{
						Code: http.StatusForbidden,
						Msg:  "您没有足够的担保金(可负额度不足)",
					})
					return
				}
			}
		}

	}

	err = TransactionDbCtrl.CarrySurety(CarrySuretyRaw.OrderID, user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "担保失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, CommunicationStructure.Message{
			Code: http.StatusOK,
			Msg:  "担保成功",
		})
		return
	}
}

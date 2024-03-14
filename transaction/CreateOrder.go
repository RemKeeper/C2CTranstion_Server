package transaction

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/SettingDbCtrl"
	"C2CTranstion_Server/TransactionDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

type CreateOrderRequest struct {
	ProductType uint8   `json:"product_type"`
	Prices      float64 `json:"prices"`
	Count       int     `json:"count"`
	IsRetail    bool    `json:"is_retail"`
}

func CreateOrder(c *gin.Context) {
	data, err := c.GetRawData()
	var CreateOrderRaw CreateOrderRequest
	err = json.Unmarshal(data, &CreateOrderRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取数据失败",
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

	var CreateOrderInsert TransactionDbCtrl.CreateOrderInsert
	CreateOrderInsert.ProductType = CreateOrderRaw.ProductType
	CreateOrderInsert.Prices = CreateOrderRaw.Prices
	CreateOrderInsert.Count = CreateOrderRaw.Count
	CreateOrderInsert.TotalPrice = CreateOrderRaw.Prices * float64(CreateOrderRaw.Count)
	CreateOrderInsert.OrderStatus = 1
	CreateOrderInsert.IsRetail = CreateOrderRaw.IsRetail

	switch CreateOrderRaw.ProductType {
	case 0:
		CreateOrderInsert.SellerID = uint32(user.UserID)
		switch {
		case CreateOrderRaw.Count%100 != 0:
			c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
				Code: http.StatusBadRequest,
				Msg:  "数量不是100的倍数",
			})
			return
		case CreateOrderRaw.Count < SettingDbCtrl.GlobalConfig.MinCoinTransaction:
			c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
				Code: http.StatusBadRequest,
				Msg:  "数量小于最小金币交易数量",
			})
			return
		case CreateOrderRaw.Count < user.CoinBalance:
			user.CoinBalance -= CreateOrderRaw.Count
			_, err = UserDbCtrl.EditUserPri(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
					Code: http.StatusInternalServerError,
					Msg:  "扣除用户余额失败",
				})
				return
			}
			CreateOrderInsert.OrderStatus = 3
		}

	case 1:
		CreateOrderInsert.PurchasersID = uint32(user.UserID)
		switch {
		case CreateOrderRaw.Count < SettingDbCtrl.GlobalConfig.MinGemTransaction:
			c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
				Code: http.StatusBadRequest,
				Msg:  fmt.Sprint("数量小于最小宝石交易数量", SettingDbCtrl.GlobalConfig.MinGemTransaction, "个"),
			})
			return
		case float64(CreateOrderRaw.Count)*CreateOrderRaw.Prices < user.Balance:
			user.Balance -= float64(CreateOrderRaw.Count) * CreateOrderRaw.Prices
			_, err = UserDbCtrl.EditUserPri(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
					Code: http.StatusInternalServerError,
					Msg:  "扣除用户余额失败",
				})
				return
			}
			CreateOrderInsert.OrderStatus = 3
		}
	default:
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "商品类型错误",
		})
	}

	err = TransactionDbCtrl.CreateOrder(CreateOrderInsert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "创建订单失败",
		})
		return
	}

	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "创建订单成功，正在寻找担保人确认",
	})

}

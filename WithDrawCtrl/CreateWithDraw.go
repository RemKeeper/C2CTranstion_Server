package WithDrawCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"C2CTranstion_Server/WithDrawDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateWithDrawRequest struct {
	ProductType uint    `json:"product_type"`
	Amount      float64 `json:"amount"`
}

func CreateWithDraw(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 获取cookie失败 " + err.Error(),
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return
	}
	var createWithDrawRequest CreateWithDrawRequest
	err = c.BindJSON(&createWithDrawRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 解析数据失败 " + err.Error(),
		})
		return
	}
	if createWithDrawRequest.Amount <= 0 {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 提现金额不能小于0 ",
		})
		return
	}
	switch createWithDrawRequest.ProductType {
	case 0:
		if user.CoinBalance < int(createWithDrawRequest.Amount) {
			c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
				Code: http.StatusBadRequest,
				Msg:  " 金币余额不足 ",
			})
			return
		}
		err := UserDbCtrl.UserTransactionEditCoinBalance(false, user.UserID, int(createWithDrawRequest.Amount))
		if err != nil {
			c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
				Code: http.StatusBadRequest,
				Msg:  " 金币余额不足 ",
			})
			return
		}
	case 1:
		if user.Balance < createWithDrawRequest.Amount {
			c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
				Code: http.StatusBadRequest,
				Msg:  " 余额不足 ",
			})
			return
		}
		err := UserDbCtrl.UserTransactionEditBalance(false, user.UserID, createWithDrawRequest.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
				Code: http.StatusBadRequest,
				Msg:  " 余额不足 ",
			})
			return
		}
	}
	err = WithDrawDbCtrl.CreateWithDraw(user, createWithDrawRequest.ProductType, createWithDrawRequest.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 创建提现失败 " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  " 创建提现申请成功 ",
	})
}

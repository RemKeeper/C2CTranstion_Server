package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/Toast"
	"C2CTranstion_Server/UserDbCtrl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransferRequest struct {
	ToUserId uint    `json:"to_user_id"`
	Amount   float64 `json:"amount"`
}

func TransferCoin(c *gin.Context) {

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

	var request TransferRequest
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败",
		})
		return
	}

	if request.ToUserId == user.UserID {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "不能给自己转账",
		})
		return

	}

	if user.CoinBalance < int(request.Amount) {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "金币余额不足",
		})
		return

	}

	ToUser, err := UserDbCtrl.QueryUserByID(request.ToUserId)
	if err != nil {
		return
	}

	if ToUser.UserName == "" {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "用户不存在",
		})
		return
	}

	err = UserDbCtrl.TransferCoin(user.UserID, request.ToUserId, int(request.Amount))
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "转账失败",
		})
		return
	}
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "转账成功",
	})
	Toast.AddMessage(Toast.Message{
		Userid: request.ToUserId,
		Msg:    fmt.Sprint("您收到了", user.UserName, "转账的", request.Amount, "金币"),
	})
}

func TransferBalance(c *gin.Context) {
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
	var request TransferRequest
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败",
		})
		return
	}
	if request.ToUserId == user.UserID {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "不能给自己转账",
		})
		return

	}

	if user.Balance < request.Amount {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "余额不足",
		})
		return
	}
	ToUser, err := UserDbCtrl.QueryUserByID(request.ToUserId)
	if ToUser.UserName == "" {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "用户不存在",
		})
		return
	}

	err = UserDbCtrl.TransferBalance(user.UserID, request.ToUserId, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "转账失败",
		})
		return
	}
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "转账成功",
	})
	Toast.AddMessage(Toast.Message{
		Userid: request.ToUserId,
		Msg:    fmt.Sprint("您收到了", user.UserName, "转账的", request.Amount, "元"),
	})
}

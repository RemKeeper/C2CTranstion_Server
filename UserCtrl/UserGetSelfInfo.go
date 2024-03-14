package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
)

type UserGetSelfResponse struct {
	UserId       uint    `json:"user_id"`
	UserName     string  `json:"user_name"`
	Balance      float64 `json:"balance"`
	CoinBalance  int     `json:"coin_balance"`
	AlipayUname  string  `json:"alipay_uname"`
	UserLastname string  `json:"user_lastname"`
	AlipayQrcode string  `json:"alipay_qrcode"`
	WxpayQrcode  string  `json:"wxpay_qrcode"`
	BankCard     string  `json:"bank_card"`
	GameId       uint    `json:"game_id"`
	Phone        uint    `json:"phone"`
}

func GetSelfInfo(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "获取cookie失败" + err.Error()})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "获取用户失败" + err.Error()})
		return
	}
	c.JSON(200, UserGetSelfResponse{
		UserId:       user.UserID,
		UserName:     user.UserName,
		Balance:      user.Balance,
		CoinBalance:  user.CoinBalance,
		AlipayUname:  user.AlipayUname,
		UserLastname: user.UserLastname,
		AlipayQrcode: user.AlipayQrcode,
		WxpayQrcode:  user.WxpayQrcode,
		BankCard:     user.BankCard,
		GameId:       user.GameID,
		Phone:        user.Phone,
	})
}

package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserGetResponse struct {
	UserId       uint   `json:"user_id"`
	UserName     string `json:"user_name"`
	AlipayUname  string `json:"alipay_uname"`
	UserLastname string `json:"user_lastname"`
	AlipayQrcode string `json:"alipay_qrcode"`
	WxpayQrcode  string `json:"wxpay_qrcode"`
	GameId       uint   `json:"game_id"`
	Phone        uint   `json:"phone"`
}

func GetUserInfo(c *gin.Context) {
	//获取get携带信息
	UidStr, exit := c.GetQuery("user_id")
	if !exit {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
		})
		return
	}
	userId, err := strconv.Atoi(UidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
		})
		return
	}
	user, err := UserDbCtrl.QueryUserByID(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "未找到用户",
		})
		return
	}

	c.JSON(http.StatusOK, UserGetResponse{
		UserId:       user.UserID,
		UserName:     user.UserName,
		AlipayUname:  user.AlipayUname,
		UserLastname: user.UserLastname,
		AlipayQrcode: user.AlipayQrcode,
		WxpayQrcode:  user.WxpayQrcode,
		GameId:       user.GameID,
		Phone:        user.Phone,
	})
}

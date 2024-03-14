package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
)

func SetSelfInfo(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "获取cookie失败" + err.Error()})
		return
	}
	var UserSet UserDbCtrl.User
	data, err := c.GetRawData()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &UserSet)
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "解析json失败" + err.Error()})
		return
	}

	err = UserDbCtrl.UserEditSelfInfo(cookie, UserDbCtrl.User{
		AlipayUname:  UserSet.AlipayUname,
		UserLastname: UserSet.UserLastname,
		AlipayQrcode: UserSet.AlipayQrcode,
		WxpayQrcode:  UserSet.WxpayQrcode,
		BankCard:     UserSet.BankCard,
		GameID:       UserSet.GameID,
		Phone:        UserSet.Phone,
	})
	//如果以上信息都不为空
	if UserSet.AlipayUname != "" && UserSet.UserLastname != "" && UserSet.AlipayQrcode != "" && UserSet.WxpayQrcode != "" && UserSet.BankCard != "" && UserSet.GameID != 0 && UserSet.Phone != 0 {
		user, err := UserDbCtrl.GetUserByCookie(cookie)
		if err != nil {
			return
		}
		user.PriLabel = 2
		UserDbCtrl.EditUserPri(user)

	}
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "修改用户信息失败" + err.Error()})
		return
	}
	c.JSON(200, CommunicationStructure.Message{Code: 200, Msg: "修改用户信息成功"})
}

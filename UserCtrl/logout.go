package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return
	}
	err = UserDbCtrl.Logout(cookie)
	if err != nil {
		return
	}
	c.SetCookie("session_id", "", -1, "/", "*", false, false)
	c.JSON(200, CommunicationStructure.Message{
		Code: 200,
		Msg:  "登出成功",
	})
}

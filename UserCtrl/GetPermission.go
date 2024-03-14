package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
)

func GetPermission(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "Bad Request",
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"level": user.PriLabel,
	})
}

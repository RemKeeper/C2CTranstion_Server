package UserRebateCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"C2CTranstion_Server/UserRebateDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserRebate(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return
	}
	userRebate, err := UserRebateDbCtrl.GetUserRebate(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  " 获取返利信息失败 " + err.Error(),
		})
		return
	}
	c.JSON(200, userRebate)
}

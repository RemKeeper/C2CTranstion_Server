package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInvitedUser(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法获取cookie",
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法获取用户信息",
		})
		return
	}
	invitedUser, err := UserDbCtrl.GetInvitedUser(user.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法获取邀请用户信息",
		})
		return
	}
	c.JSON(http.StatusOK, invitedUser)
}

func GetInvitedUserRecursive(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法获取cookie",
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法获取用户信息",
		})
		return
	}
	var invitedMap = make(map[uint][]uint)
	invitedUser, err := UserDbCtrl.GetInvitedUser(user.UserID)
	for _, u := range invitedUser {
		invitedUser, err = UserDbCtrl.GetInvitedUser(u)
		if err != nil {
			continue
		}
		invitedMap[u] = append(invitedMap[u], invitedUser...)
	}
	c.JSON(http.StatusOK, invitedMap)
}

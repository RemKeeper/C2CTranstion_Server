package Middleware

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var PagePermissionHierarchy = map[string]uint{
	"/register": 0,
	"/login":    0,

	"/logout": 0,

	"/image":  1,
	"/upload": 2,

	"/transaction": 2,

	"/admin": 5,

	"/user": 1,

	"/surety": 3,
}

func CheckPri(c *gin.Context) {

	path := "/" + strings.Split(c.Request.URL.Path, "/")[1]
	fmt.Println(path)
	switch path {
	case "/register":
		c.Next()
	case "/login":
		c.Next()
	default:
		cookie, err := c.Cookie("session_id")
		if err != nil {
			//c.Redirect(302, "/login")
			c.JSON(http.StatusForbidden, CommunicationStructure.Message{
				Code: http.StatusForbidden,
				Msg:  "权限不足，请检查是否登录",
			})
			c.Abort()
			return
		}
		user, err := UserDbCtrl.GetUserByCookie(cookie)
		if err != nil {
			c.SetCookie("session_id", "", -1, "/", "*", false, false)
			c.JSON(http.StatusForbidden, CommunicationStructure.Message{
				Code: http.StatusForbidden,
				Msg:  "权限不足，请检查是否登录",
			})
			c.Abort()
			return
		}
		if user.IsFreeze {
			c.JSON(http.StatusForbidden, CommunicationStructure.Message{
				Code: http.StatusForbidden,
				Msg:  "账户已被冻结，请联系管理员",
			})
			c.SetCookie("session_id", "", -1, "/", "*", false, false)
			c.Abort()
			return
		}
		if user.PriLabel < PagePermissionHierarchy[path] {
			c.JSON(http.StatusForbidden, CommunicationStructure.Message{
				Code: http.StatusForbidden,
				Msg:  "权限不足，请检查您的个人信息是否完整或者是否登录",
			})
			c.Abort()
			return
		}
	}
	c.Next()
}

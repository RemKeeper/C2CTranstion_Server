package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"C2CTranstion_Server/Utils"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func Login(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法读取请求数据",
		})
		return
	}
	var UserData AccountMessages
	err = json.Unmarshal(data, &UserData)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法解析请求数据",
		})
		return
	}
	user, QueryResult := UserDbCtrl.QueryUserExit(UserDbCtrl.User{
		UserName:   UserData.UserName,
		PwdSummary: UserData.PwdSummary,
	})
	if QueryResult == false {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "用户名或密码错误",
		})
		return
	}
	if user.IsFreeze == true {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "该用户已被冻结",
		})
		return

	}
	cookie := Utils.GenCookie(UserData.UserName, UserData.PwdSummary)
	err = UserDbCtrl.AddCookie(UserDbCtrl.Cookie{
		UserID: user.UserID,
		Cookie: cookie,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "生成Cookie失败",
		})
		return
	}
	c.SetCookie("session_id", cookie, 259200, "/", "", false, false)
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "登录成功",
	})
	return

}

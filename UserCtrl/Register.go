package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"C2CTranstion_Server/Utils"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"log"
	"net/http"
)

type AccountMessages struct {
	UserName       string `json:"user_name"`
	PwdSummary     string `json:"pwd_summary"`
	InvitationCode uint   `json:"invitation_code"`
}

func Register(c *gin.Context) {
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

	//空值判断
	if UserData.UserName == "" || UserData.PwdSummary == "" {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "用户名或密码为空",
		})
		return
	}

	if !Utils.IsValidInput(UserData.UserName) == true {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "用户名不合法",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "无法解析请求数据",
		})
		return
	}
	log.Println(UserData)
	userId, err := UserDbCtrl.AddUser(UserDbCtrl.User{
		UserName:   UserData.UserName,
		PwdSummary: UserData.PwdSummary,
		InviterId:  UserData.InvitationCode,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "用户名已存在",
		})
		return
	}
	cookie := Utils.GenCookie(UserData.UserName, UserData.PwdSummary)
	err = UserDbCtrl.AddCookie(UserDbCtrl.Cookie{
		UserID: userId,
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
		Msg:  "注册成功",
	})
}

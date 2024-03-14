package SettingCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/SettingDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSetting(c *gin.Context) {
	config, err := SettingDbCtrl.GetSetting()
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "获取配置失败 " + err.Error(),
		})
		return
	}
	c.JSON(200, config)
}

package SettingCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/SettingDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func SaveSetting(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "获取数据失败 " + err.Error(),
		})
		return
	}
	var config SettingDbCtrl.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "解析数据失败 " + err.Error(),
		})
		return
	}
	err = SettingDbCtrl.SaveSetting(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "保存配置失败 " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "保存配置成功",
	})
}

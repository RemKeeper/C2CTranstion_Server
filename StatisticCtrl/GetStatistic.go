package StatisticCtrl

import (
	"C2CTranstion_Server/StatisticDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatistic(c *gin.Context) {
	Statistic := StatisticDbCtrl.GetStatistic()
	c.JSON(http.StatusOK, Statistic)
}

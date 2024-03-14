package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type IPData struct {
	Requests  int
	Timestamp int64
}

var IpList = make(map[string]IPData)
var BanList = make(map[string]bool)
var rwMutex = &sync.RWMutex{}

func CheckIp(c *gin.Context) {
	ip := c.ClientIP()
	NowUnix := time.Now().Unix()

	rwMutex.RLock()
	requestData, ok := IpList[ip]
	banned := BanList[ip]
	rwMutex.RUnlock()

	if banned {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "message": "请求次数过多 IP 已被封禁"})
		c.Abort()
		return
	}

	if ok {
		// IP 存在，增加请求次数
		requestData.Requests++
		if NowUnix-requestData.Timestamp > 60 {
			// 超过 60 秒，重置请求次数
			requestData.Requests = 1
			requestData.Timestamp = NowUnix
		}
		if requestData.Requests > 100 {
			rwMutex.Lock()
			BanList[ip] = true
			rwMutex.Unlock()
			c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "message": "请求次数过多 IP 已被封禁"})
			c.Abort()
			return
		}
	} else {
		// IP 不存在，初始化数据
		requestData = IPData{
			Requests:  1,
			Timestamp: time.Now().Unix(),
		}
	}

	rwMutex.Lock()
	IpList[ip] = requestData
	rwMutex.Unlock()

	c.Next()

	if len(IpList) > 5000 {
		rwMutex.Lock()
		for s, data := range IpList {
			if NowUnix-data.Timestamp > 86400 {
				delete(IpList, s)
			}
		}
		rwMutex.Unlock()
	}
}

func CheckRequestSize(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1024*1024*10)
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "请求过大"})
		c.Abort()
		return
	}
	c.Next()
}

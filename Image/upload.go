package Image

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/ImageDbCtrl"
	"C2CTranstion_Server/OrderDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {
	c.Request.ContentLength = 2 * 1024 * 1024
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	file := form.File["file"][0]
	orderID := form.Value["order_id"][0]
	parseUint, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return
	}
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return
	}

	Order, err := OrderDbCtrl.GetOrderById(uint32(parseUint))
	if err != nil {
		return
	}

	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "上传失败 " + err.Error(),
		})
		return
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{
			Code: 500,
			Msg:  "无法打开文件",
		})
		return
	}
	defer openedFile.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, openedFile); err != nil {
		c.JSON(500, CommunicationStructure.Message{
			Code: 500,
			Msg:  "无法读取文件",
		})
		return
	}

	md5String := hex.EncodeToString(hash.Sum(nil))

	switch {
	case uint(Order.PurchasersID) == user.UserID:
		err := ImageDbCtrl.CreateOrderImg(ImageDbCtrl.OrderImg{
			OrderID:       int(parseUint),
			PurchasersImg: md5String,
			Time:          time.Now(),
		})
		if err != nil {
			c.JSON(500, CommunicationStructure.Message{
				Code: 500,
				Msg:  "上传失败",
			})
			return
		}
	case uint(Order.SellerID) == user.UserID:
		err := ImageDbCtrl.CreateOrderImg(ImageDbCtrl.OrderImg{
			OrderID:   int(parseUint),
			SellerImg: md5String,
			Time:      time.Now(),
		})
		if err != nil {
			c.JSON(500, CommunicationStructure.Message{
				Code: 500,
				Msg:  "上传失败",
			})
			return
		}
	default:
		c.JSON(500, CommunicationStructure.Message{
			Code: 500,
			Msg:  "上传失败，请确认订单ID",
		})
		return
	}

	FileName := md5String + ".jpg"
	err = c.SaveUploadedFile(file, "./images/"+FileName)
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{
			Code: 500,
			Msg:  "上传失败",
		})
		return
	}
	err = OrderDbCtrl.UpdateDispatchStatus(Order.OrderID)
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{
			Code: 500,
			Msg:  "上传失败",
		})
		return
	}
	c.JSON(200, CommunicationStructure.Message{
		Code: 200,
		Msg:  "上传成功",
	})
}

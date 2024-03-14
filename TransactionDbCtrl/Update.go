package TransactionDbCtrl

import (
	"C2CTranstion_Server/DividendPercentageDbCtrl"
	"C2CTranstion_Server/OrderDbCtrl"
	"C2CTranstion_Server/SettingDbCtrl"
	"C2CTranstion_Server/Toast"
	"C2CTranstion_Server/UserDbCtrl"
	"fmt"
)

func CarrySurety(orderId uint32, userId uint) error {
	var order OrderDbCtrl.Order
	err := Db.Model(OrderDbCtrl.Order{}).Where("order_id = ? AND surety_id IS NULL", orderId).Update("surety_id", userId).First(&order).Error
	if err != nil {
		return err
	}
	if order.ProductType == 0 {

		err = UserDbCtrl.UserTransactionEditCoinBalance(false, uint(order.SellerID), order.Count)
		if err != nil {
			return err
		}
		Toast.AddMessage(Toast.Message{
			Userid: uint(order.PurchasersID),
			Msg:    fmt.Sprint("您的订单:", order.OrderID, " 已被担保人:", userId, "接单,请尽快付款"),
		})
	} else {
		err = UserDbCtrl.UserTransactionEditBalance(false, uint(order.PurchasersID), order.TotalPrice)
		if err != nil {
			return err
		}
		Toast.AddMessage(Toast.Message{
			Userid: uint(order.SellerID),
			Msg:    fmt.Sprint("您的订单:", order.OrderID, " 已被担保人:", userId, "接单,请尽快付款"),
		})
	}

	return nil
}

func ConfirmSurety(orderId uint32, userId uint) {
	var order OrderDbCtrl.Order
	Db.Model(OrderDbCtrl.Order{}).Where("order_id = ? AND surety_id = ?", orderId, userId).Update("order_status", 0).First(&order)

	if order.ProductType == 0 {
		_ = UserDbCtrl.UserTransactionEditCoinBalance(true, uint(order.SuretyID), int(float64(order.Count)*SettingDbCtrl.GlobalConfig.BountyPercentage))
		Toast.AddMessage(Toast.Message{
			Userid: uint(order.SellerID),
			Msg:    fmt.Sprint("您的订单:", order.OrderID, " 已被担保人:", order.SuretyID, "确认上架"),
		})
	} else {
		_ = UserDbCtrl.UserTransactionEditBalance(true, uint(order.SuretyID), order.TotalPrice*SettingDbCtrl.GlobalConfig.BountyPercentage)
		Toast.AddMessage(Toast.Message{
			Userid: uint(order.PurchasersID),
			Msg:    fmt.Sprint("您的订单:", order.OrderID, " 已被担保人:", order.SuretyID, "确认上架"),
		})
	}

}

func CarryTransaction(order OrderDbCtrl.Order, user UserDbCtrl.User) error {
	switch order.ProductType {
	case 0:
		//判断余额是否充足
		if user.Balance >= order.TotalPrice {
			//减少买家余额
			err := UserDbCtrl.UserTransactionEditBalance(false, user.UserID, order.TotalPrice)
			if err != nil {
				return err
			}
			//增加卖家余额
			err = UserDbCtrl.UserTransactionEditBalance(true, uint(order.SellerID), order.TotalPrice)
			if err != nil {
				return err
			}
			//增加买家金币余额
			err = UserDbCtrl.UserTransactionEditCoinBalance(true, user.UserID, order.Count)
			if err != nil {
				return err
			}
			//更改订单状态
			err = Db.Model(OrderDbCtrl.Order{}).Where("order_id = ?", order.OrderID).Update("order_status", 2).Update("purchasers_id", user.UserID).Error
			if err != nil {
				return err
			}
			//通知卖家
			Toast.AddMessage(Toast.Message{
				Userid: uint(order.SellerID),
				Msg:    fmt.Sprint("您的订单:", order.OrderID, " 已被买家:", user.UserID, "余额交易成功"),
			})
			return nil
		} else {
			err := Db.Model(OrderDbCtrl.Order{}).Where("order_id = ?", order.OrderID).Update("order_status", 5).Update("purchasers_id", user.UserID).Error
			if err != nil {
				return err
			}
			//通知卖家
			Toast.AddMessage(Toast.Message{
				Userid: uint(order.SellerID),
				Msg:    fmt.Sprint("您的订单:", order.OrderID, " \n已被买家:", user.UserID, "\n拍下，等待付款"),
			})
			//通知担保人
			Toast.AddMessage(Toast.Message{
				Userid: uint(order.SuretyID),
				Msg:    fmt.Sprint("您担保的订单:", order.OrderID, " \n已被买家:", user.UserID, "\n拍下，等待付款"),
			})
		}
	case 1:
		err := Db.Model(OrderDbCtrl.Order{}).Where("order_id = ?", order.OrderID).Update("order_status", 5).Update("seller_id", user.UserID).Error
		if err != nil {
			return err
		}
		//通知卖家
		Toast.AddMessage(Toast.Message{
			Userid: uint(order.SellerID),
			Msg:    fmt.Sprint("您的订单:", order.OrderID, " \n已被卖家:", user.UserID, "\n拍下，等待付款"),
		})
		//通知担保人
		Toast.AddMessage(Toast.Message{
			Userid: uint(order.SuretyID),
			Msg:    fmt.Sprint("您担保的订单:", order.OrderID, " \n已被卖家:", user.UserID, "\n拍下，等待付款"),
		})
	}
	return nil
}

func ConfirmTransaction(orderId uint32, userId uint) {
	var FindOrder OrderDbCtrl.Order
	Db.Where("order_id = ?", orderId).First(&FindOrder)

	if FindOrder.OrderStatus == 5 {
		switch FindOrder.ProductType {
		case 0:
			Db.Where("order_id = ? AND purchasers_id = ?", orderId, userId).Update("order_status", 2)
			Toast.AddMessage(Toast.Message{
				Userid: uint(FindOrder.PurchasersID),
				Msg:    fmt.Sprint("您的订单:", FindOrder.OrderID, " 已被卖家:", userId, "确认"),
			})
		case 1:
			Db.Where("order_id = ? AND seller_id = ?", orderId, userId).Update("order_status", 2)
			Toast.AddMessage(Toast.Message{
				Userid: uint(FindOrder.SellerID),
				Msg:    fmt.Sprint("您的订单:", FindOrder.OrderID, " 已被买家:", userId, "确认"),
			})
		}
		_ = DividendPercentageDbCtrl.GetDividendPercentage(FindOrder)

	}
}

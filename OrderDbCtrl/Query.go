package OrderDbCtrl

import (
	"C2CTranstion_Server/UserDbCtrl"
)

func GetAllOrder(sortStr string) ([]Order, error) {
	var OrderList []Order
	err := Db.Order(sortStr).Find(&OrderList).Error
	if err != nil {
		return nil, err
	}
	return OrderList, nil
}

func GetUnfinished(sortStr string) ([]Order, error) {
	var OrderList []Order
	err := Db.Where("order_status = ?", 2).Order(sortStr).Find(&OrderList).Error
	if err != nil {
		return nil, err
	}
	return OrderList, nil
}

func UserGetAllOrder(cookie string) ([]Order, error) {
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return nil, err
	}
	var OrderList []Order
	err = Db.Where("purchasers_id = ? OR seller_id = ?", user.UserID, user.UserID).Find(&OrderList).Error
	if err != nil {
		return nil, err
	}
	return OrderList, nil
}

func UserGetWaitOrder(cookie string) ([]Order, error) {
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return nil, err
	}
	var OrderList []Order
	err = Db.Where("purchasers_id = ? OR seller_id = ?", user.UserID, user.UserID).Where("order_status = ?", 1).Find(&OrderList).Error
	if err != nil {
		return nil, err
	}
	return OrderList, nil
}

func SuretyGetWaitOrder() ([]Order, error) {
	var OrderList []Order
	err := Db.Where("order_status = ?", 1).Find(&OrderList).Error
	if err != nil {
		return nil, err
	}
	return OrderList, nil
}

func UserGetFreezeOrder(cookie string) ([]Order, error) {
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return nil, err
	}
	var OrderList []Order
	err = Db.Where("purchasers_id = ? OR seller_id = ?", user.UserID, user.UserID).Where("order_status = ?", 4).Find(&OrderList).Error
	if err != nil {
		return nil, err
	}
	return OrderList, nil
}

func SuretyGetSelfOrder(cookie string) ([]Order, error) {
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		return nil, err
	}
	var OrderList []Order
	err = Db.Where("seller_id = ?", user.UserID).Find(&OrderList).Error
	if err != nil {
		return nil, err
	}

	return OrderList, nil
}

func GetOrderById(id uint32) (Order, error) {
	var order Order
	err := Db.Where("order_id = ?", id).First(&order).Error
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

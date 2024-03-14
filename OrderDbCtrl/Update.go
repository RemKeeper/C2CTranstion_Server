package OrderDbCtrl

import "gorm.io/gorm"

func FreezeOrder(orderId int) error {
	err := Db.Model(&Order{}).Where("order_id = ?", orderId).Update("order_status", 4).Update("freeze_reasons", "管理员主动冻结").Error
	if err != nil {
		return err
	}
	return nil
}

func UnfreezeOrder(orderId int) error {
	err := Db.Model(&Order{}).Where("order_id = ?", orderId).Update("order_status", 0).Update("freeze_reasons", "").Error
	if err != nil {
		return err
	}
	return nil
}

func UserFreezeOrder(orderId, userid int) error {
	err := Db.Model(&Order{}).Where("order_id = ?", orderId).Where("purchasers_id = ? OR seller_id = ?", userid, userid).Update("order_status", 4).Update("freeze_reasons", "用户纠纷").Error
	if err != nil {
		return err
	}
	return nil
}

func TransactionSuccess(orderId int) error {
	err := Db.Model(&Order{}).Where("order_id = ?", orderId).Update("order_status", 2).Error
	if err != nil {
		return err
	}
	return nil
}

func TransactionEditOrderCount(IsAdd bool, Count int, orderId int) error {
	if IsAdd {
		err := Db.Model(&Order{}).Where("order_id = ?", orderId).Update("count", gorm.Expr("count + ", Count)).Error
		if err != nil {
			return err
		}
		return nil
	}
	err := Db.Model(&Order{}).Where("order_id = ?", orderId).Update("count", gorm.Expr("count - ", Count)).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateDispatchStatus(orderId uint32) error {
	err := Db.Model(&Order{}).Where("order_id = ?", orderId).Update("Is_Dispatched", 1).Error
	return err
}

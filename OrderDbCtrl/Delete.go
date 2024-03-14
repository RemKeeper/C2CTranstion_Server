package OrderDbCtrl

func DeleteOrder(orderid int) error {

	err := Db.Where("order_id = ?", orderid).Delete(&Order{}).Error
	if err != nil {
		return err
	}
	return nil
}

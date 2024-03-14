package SettingDbCtrl

func SaveSetting(config Config) error {
	err := Db.Where("edit_here = ?", 0).Save(&config).Error
	return err

}

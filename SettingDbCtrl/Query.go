package SettingDbCtrl

func GetSetting() (Config, error) {
	var config Config
	err := Db.First(&config).Error
	return config, err
}

package UserDbCtrl

func DeleteUser(user User) (uint, error) {
	if err := Db.Where("user_id= ? ", user.UserID).Delete(&user).Error; err != nil {
		return 0, err
	}
	return user.UserID, nil

}

func Logout(cookie string) error {
	return Db.Model(Cookie{}).Where("cookie = ?", cookie).Update("cookie", "").Error
}

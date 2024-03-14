package UserDbCtrl

import (
	"errors"
	"gorm.io/gorm"
)

func GetAllUser() ([]User, error) {
	var users []User
	err := Db.Order("user_id").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func QueryUserByID(id uint) (User, error) {
	var user User
	err := Db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func QueryUserExit(user User) (User, bool) {
	if err := Db.Where("user_name = ? AND pwd_summary = ?", user.UserName, user.PwdSummary).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, false
		}
		// handle other types of errors as necessary
	}
	return user, true
}

func GetUserByCookie(cookie string) (User, error) {
	var cookieObj Cookie
	err := Db.Where("cookie = ?", cookie).First(&cookieObj).Error
	if err != nil {
		return User{}, err
	}
	var user User
	err = Db.Where("user_id = ?", cookieObj.UserID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Db.Delete(&cookieObj)
		}
		return User{}, err
	}

	return user, nil
}

func GetInvitedUser(userId uint) ([]uint, error) {
	var userIds []uint
	err := Db.Model(&User{}).Where("inviter_id = ?", userId).Pluck("user_id", &userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

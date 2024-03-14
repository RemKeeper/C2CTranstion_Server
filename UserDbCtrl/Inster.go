package UserDbCtrl

import (
	"time"
)

func AddUser(user User) (uint, error) {
	NewUser := User{
		UserName:   user.UserName,
		PwdSummary: user.PwdSummary,
		InviterId:  user.InviterId,
		PriLabel:   1,
	}
	result := Db.Create(&NewUser).Last(&NewUser)
	if result.Error != nil {
		return 0, result.Error
	}
	return NewUser.UserID, nil
}

func AddCookie(cookie Cookie) error {
	NewCookie := Cookie{
		UserID:     cookie.UserID,
		Cookie:     cookie.Cookie,
		CreateTime: int(time.Now().Unix()),
	}
	result := Db.Create(&NewCookie)
	if result.Error != nil {
		err := Db.Model(&Cookie{}).Where(&Cookie{UserID: cookie.UserID}).Updates(&NewCookie).Error
		if err != nil {
			return err
		}
	}
	return nil
}

package UserDbCtrl

type Cookie struct {
	UserID     uint   `gorm:"primary_key;column:user_id;type:int(8) unsigned zerofill;not null;unique"`
	Cookie     string `gorm:"primary_key;column:cookie;type:varchar(255);not null"`
	CreateTime int    `gorm:"column:CreateTime;type:int;not null"`
}

func (Cookie) TableName() string {
	return "cookie"
}

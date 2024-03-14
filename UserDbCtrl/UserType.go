package UserDbCtrl

type User struct {
	UserID           uint    `gorm:"primary_key;AUTO_INCREMENT;column:user_id;type:int(8) unsigned zerofill;not null;" json:"user_id"`
	UserName         string  `gorm:"unique;column:user_name;type:char(30);size:30;not null;" json:"user_name"`
	InviterId        uint    `gorm:"column:inviter_id;type:int unsigned;" json:"inviter_id"`
	PwdSummary       string  `gorm:"column:pwd_summary;type:char(255);size:255;not null;" json:"pwd_summary"`
	PriLabel         uint    `gorm:"column:pri_label;type:int unsigned;not null;" json:"pri_label"`
	IsFreeze         bool    `gorm:"column:is_freeze;type:tinyint;default:0;not null;" json:"is_freeze"`
	Balance          float64 `gorm:"column:balance;type:float;not null;" json:"balance"`
	CoinBalanceQuota int     `gorm:"column:coin_balance_quota;type:int;not null;" json:"coin_balance_quota"`
	BalanceQuota     float64 `gorm:"column:balance_quota;type:float;not null;" json:"balance_quota"`
	CoinBalance      int     `gorm:"column:coin_balance;type:int;not null;" json:"coin_balance"`
	AlipayUname      string  `gorm:"column:alipay_uname;type:char(30);size:30;" json:"alipay_uname"`
	UserLastname     string  `gorm:"column:user_lastname;type:char(8);size:8;" json:"user_lastname"`
	AlipayQrcode     string  `gorm:"column:alipay_qrcode;type:char(255);size:255;" json:"alipay_qrcode"`
	WxpayQrcode      string  `gorm:"column:wxpay_qrcode;type:char(255);size:255;" json:"wxpay_qrcode"`
	BankCard         string  `gorm:"column:bank_card;type:char(19);size:19;" json:"bank_card"`
	GameID           uint    `gorm:"column:game_id;type:int unsigned;" json:"game_id"`
	Phone            uint    `gorm:"column:phone;type:int unsigned;" json:"phone"`
}

func (User) TableName() string {
	return "user"
}

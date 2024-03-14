package DbInit

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func init() {
	var err error
	//Db, err = ConnectDb("c2ctranstion_user:user1@tcp(127.0.0.1:3306)/c2ctranstion?charset=utf8mb4&parseTime=True&loc=Local")
	//Db, err = ConnectDb("root:root@tcp(127.0.0.1:3306)/c2ctranstion?charset=utf8mb4&parseTime=True&loc=Local")
	Db, err = ConnectDb("root:root@tcp(127.0.0.1:3306)/c2ctranstion?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("数据库连接失败")
		os.Exit(404)
	}
}

func ConnectDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		panic("数据库连接失败")
	}
	sqlDb, err := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db, err
}

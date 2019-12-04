package constant

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB

func Dbinit() {
	var err error
	Db, err = gorm.Open("mysql", "mt:mt123qwe@tcp(10.151.11.3:3306)/mt_merchant")
	if err != nil {
		panic("数据库连接失败")
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxIdleConns(5)
	Db.DB().SetConnMaxLifetime(time.Hour)
	Db.LogMode(true)

}

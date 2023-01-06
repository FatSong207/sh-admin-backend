package initializer

import (
	"SH-admin/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitGorm() {
	m := global.Config.Mysql
	s := "%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local"
	var dsn = fmt.Sprintf(s, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	global.Db = db
}

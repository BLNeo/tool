package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"xorm.io/xorm"
)

type Instance struct {
	User         string `toml:"user"`
	Password     string `toml:"password"`
	Host         string `toml:"host"`
	DatabaseName string `toml:"databaseName"`
	Charset      string `toml:"charset"`
	LogShow      bool   `toml:"log_show"`
}

// InitEngine xorm的连接方法
func InitEngine(ins *Instance) (*xorm.Engine, error) {
	dbUrl := fmt.Sprintf("%v:%v@(%v)/%v?charset=%v", ins.User, ins.Password, ins.Host, ins.DatabaseName, ins.Charset)
	dbEngine, err := xorm.NewEngine("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	dbEngine.DatabaseTZ = time.Local
	dbEngine.TZLocation = time.Local
	// 设置连接池最大打开连接数（服务器核心数*2+有效磁盘数）
	dbEngine.SetMaxOpenConns(50)
	// sql显示输出
	dbEngine.ShowSQL(ins.LogShow)
	// 设置连接最大存活时长 必须小于mysql的wait_timeout
	dbEngine.SetConnMaxLifetime(60 * time.Second)

	return dbEngine, err
}

func InitDB(ins *Instance) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		ins.User, ins.Password, ins.Host, ins.DatabaseName, ins.Charset)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        255, // string 类型字段的默认长度
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项 不会在尾部加"s"
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns设置空闲状态下的最大连接数
	sqlDB.SetMaxIdleConns(5)
	// SetMaxOpenConns设置数据库打开的最大连接数
	sqlDB.SetMaxOpenConns(10)
	// ping
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

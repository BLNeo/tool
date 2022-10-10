package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func InitEngine(ins *Instance) (*xorm.Engine, error) {
	dbUrl := fmt.Sprintf("%v:%v@(%v)/%v?charset=%v", ins.User, ins.Password, ins.Host, ins.Database, ins.Charset)
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

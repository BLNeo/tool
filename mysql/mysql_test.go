package mysql

import (
	"testing"
)

func TestMysql(t *testing.T) {
	ins := &Instance{
		User:     "root",
		Password: "123456",
		Host:     "127.0.0.1:3306",
		Database: "test_mysql",
		Charset:  "utf8mb4",
		LogShow:  false,
	}
	db, err := InitEngine(ins)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
}

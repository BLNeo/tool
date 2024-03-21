package mysql

import (
	"testing"
)

func TestMysql(t *testing.T) {
	ins := &Instance{
		User:         "root",
		Password:     "123456",
		Host:         "127.0.0.1:3306",
		DatabaseName: "test",
		Charset:      "utf8mb4",
		LogShow:      false,
	}
	// xorm
	xDb, err := InitEngine(ins)
	if err != nil {
		t.Fatal(err)
	}
	err = xDb.Ping()
	if err != nil {
		t.Fatal(err)
	}

	// gorm
	_, err = InitDB(ins)
	if err != nil {
		t.Fatal(err)
	}

}

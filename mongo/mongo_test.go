package mongo

import "testing"

func TestMongo(t *testing.T) {
	ins := &Instance{
		User:      "",
		Password:  "",
		DbName:    "test_mongo",
		Addresses: []string{"127.0.0.1:27017"},
		Option:    nil,
	}
	_, err := InitEngine(ins)
	if err != nil {
		t.Fatal(err)
	}
}

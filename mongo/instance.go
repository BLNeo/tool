package mongo

import "errors"

type Instance struct {
	User      string            `toml:"user"`
	Password  string            `toml:"password"`
	DbName    string            `toml:"db_name"`
	Addresses []string          `toml:"addresses"`
	Option    map[string]string `toml:"option"`
}

func (ins *Instance) String() (string, error) {
	str := "mongodb://"

	// "" ,""
	if ins.User != "" && ins.Password != "" {
		str += ins.User + ":" + ins.Password + "@"
	}

	addr := ""
	if len(ins.Addresses) == 0 {
		return "", errors.New("mongodb : no set address")
	}
	for k, v := range ins.Addresses {
		if k == 0 {
			addr += v // localhost1
			continue
		}
		addr += "," + v // localhost1,localhost2
	}
	str += addr + "/"

	if ins.DbName != "" {
		str += ins.DbName
	}
	str += "?"

	for k, v := range ins.Option {
		// todo 这里最好有一个验证器，去过滤掉不合规的参数
		temStr := k + "=" + v + "&"
		str += temStr
	}
	return str, nil
}

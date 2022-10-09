package mongo

type Instance struct {
	User      string   `toml:"user"`
	Password  string   `toml:"password"`
	DbName    string   `toml:"db_name"`
	Addresses []string `toml:"addresses"`
}

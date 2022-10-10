package mysql

type Instance struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Charset  string `toml:"charset"`
	LogShow  bool   `toml:"log_show"`
}

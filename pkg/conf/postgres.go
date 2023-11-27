package conf

type Postgres struct {
	Address  string `yaml:"address"`
	Port     string `json:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

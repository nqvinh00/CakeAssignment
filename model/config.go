package model

type Config struct {
	HTTP HTTP     `yaml:"http"`
	DB   Database `yaml:"database"`
}

type HTTP struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	SecretKey string `yaml:"secret_key"`
}

type Database struct {
	DB              string `yaml:"db"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Host            string `yaml:"host"`
	Args            string `yaml:"args"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idel_conns"`
	MaxConnLifeTime int    `yaml:"max_conn_life_time"`
}

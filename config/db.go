package config

type MySQLConf struct {
	Host          string
	Port          string
	Username      string
	Password      string
	DefaultDB     string
	MaxConnection uint
}

type MongoDBConf struct {
	Host     string
	Port     string
	Username string
	Password string
}

type RedisConf struct {
	Host     string
	Port     string
	Username string
	Password string
}

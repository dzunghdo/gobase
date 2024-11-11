package config

import "gobase/config/flags"

type AppConfig struct {
	App       AppInfo
	Security  SecurityConf
	MySQL     MySQLConf
	MongoDB   MongoDBConf
	Redis     RedisConf
	Datadog   DataDogConf
	Websocket WebSocketConf
}

var appConfig AppConfig

type AppInfo struct {
	Name string
	Port string
	Env  string
}

type WebSocketConf struct {
	MaxWSConn int
}

type SecurityConf struct {
	JWTSecret []byte
}

type DataDogConf struct {
	DDAPIKey string
}

func LoadConfig() {
	flags.LoadFlags()
	appConfig = AppConfig{
		App: AppInfo{
			Name: flags.AppName,
			Port: flags.Port,
			Env:  flags.Env,
		},
		Security: SecurityConf{
			JWTSecret: []byte(flags.JWTSecret),
		},
		MySQL: MySQLConf{
			Host:          flags.MySQLHost,
			Port:          flags.MySQLPort,
			Username:      flags.MySQLUser,
			Password:      flags.MySQLPass,
			DefaultDB:     flags.MySQLDB,
			MaxConnection: 10,
		},
		Redis: RedisConf{
			Host:     flags.RedisHost,
			Port:     flags.RedisPort,
			Username: "",
			Password: "",
		},
		Datadog: DataDogConf{
			DDAPIKey: flags.DDAPIKey,
		},
		Websocket: WebSocketConf{
			MaxWSConn: flags.MaxWebsocketConn,
		},
	}
}

func GetConfig() AppConfig {
	return appConfig
}

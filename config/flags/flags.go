package flags

import "flag"

// AppInfo flags
var (
	AppName = *flag.String("APP_NAME", "go_base", "App name")
	Port    = *flag.String("PORT", "8888", "App port")
	Env     = *flag.String("ENV", "dev", "Environment")
)

// MySQL flags
var (
	MySQLHost = *flag.String("MYSQL_HOST", "localhost", "MySQL host")
	MySQLPort = *flag.String("MYSQL_PORT", "13306", "MySQL port")
	MySQLUser = *flag.String("MYSQL_USER", "root", "MySQL user")
	MySQLPass = *flag.String("MYSQL_PASS", "root", "MySQL password")
	MySQLDB   = *flag.String("MYSQL_DB", "go_base", "MySQL database")
)

// Redis flags
var (
	RedisHost = *flag.String("REDIS_HOST", "localhost", "Redis host")
	RedisPort = *flag.String("REDIS_PORT", "16379", "Redis port")
)

var (
	MaxWebsocketConn = *flag.Int("MAX_WS_CONN", 1000, "Max number of ws connections")
)

// Security flags
var (
	JWTSecret = *flag.String("JWT_SECRET", "your_jwt_secret_key", "JWT secret key")
)

// Datadog
var (
	DDAPIKey = *flag.String("DD_API_KEY", "", "Datadog API key")
)

func LoadFlags() {
	flag.Parse()
}

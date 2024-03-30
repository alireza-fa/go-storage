package constants

// general
const (
	PORT        = "PORT"
	Development = "DEVELOPMENT"
)

// logger
const (
	LogLevel   = "LOG_LEVEL"
	LoggerName = "LOGGER_NAME"

	ZapEncoding = "ZAP_ENCODING"
	ZapFilePath = "ZAP_FILE_PATH"

	SeqApiKey  = "SEQ_API_KEY"
	SeqBaseUrl = "SEQ_BASE_URL"
	SeqPort    = "SEQ_PORT"
)

// Database
const (
	DbHost      = "DB_HOST"
	DbHostDebug = "DB_HOST_DEBUG"
	DbPort      = "DB_PORT"
	DbUsername  = "DB_USERNAME"
	DbPassword  = "DB_PASSWORD"
	DbDatabase  = "DB_DATABASE"
)

// Redis
const (
	RedisHost          = "REDIS_HOST"
	RedisDebugHost     = "REDIS_DEBUG_HOST"
	RedisPort          = "REDIS_PORT"
	RedisPassword      = "REDIS_PASSWORD"
	RedisDb            = "REDIS_DB"
	DialTimeout        = "DIAL_TIMEOUT"
	ReadTimeout        = "READ_TIMEOUT"
	WriteTimeout       = "WRITE_TIMEOUT"
	PoolSize           = "POOL_SIZE"
	PoolTimeout        = "POOL_TIMEOUT"
	IdleCheckFrequency = "IDLE_CHECK_FREQUENCY"
)

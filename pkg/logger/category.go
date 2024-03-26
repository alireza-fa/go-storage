package logger

type Category string
type SubCategory string
type ExtraKey string

const (
	General    Category = "General"
	RDBMS      Category = "Postgres"
	Server     Category = "Server"
	Validation Category = "Validation"
	Redis      Category = "Redis"
)

const (
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Insert    SubCategory = "Insert"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Rollback  SubCategory = "Rollback"
	Commit    SubCategory = "Commit"

	BodyParser SubCategory = "BodyParser"

	RedisGet SubCategory = "RedisGet"
	RedisSet SubCategory = "RedisSet"
)

const (
	AppName ExtraKey = "AppName"
	Signal  ExtraKey = "Signal"
	Error   ExtraKey = "Error"
)

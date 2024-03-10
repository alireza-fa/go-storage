package logger

type Category string
type SubCategory string
type ExtraKey string

const (
	General  Category = "General"
	Postgres Category = "Postgres"
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
)

const (
	AppName ExtraKey = "AppName"
)

package logger

type Logger interface {
	Init()

	Debug(cat Category, Sub SubCategory, message string, extra map[ExtraKey]interface{})

	Info(cat Category, Sub SubCategory, message string, extra map[ExtraKey]interface{})

	Warn(cat Category, Sub SubCategory, message string, extra map[ExtraKey]interface{})

	Error(cat Category, Sub SubCategory, message string, extra map[ExtraKey]interface{})

	Fatal(cat Category, Sub SubCategory, message string, extra map[ExtraKey]interface{})
}

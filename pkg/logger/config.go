package logger

type Config struct {
	Logger      string
	Development bool
	Encoding    string
	Level       string
	FilePath    string
	Seq         struct {
		ApiKey  string
		BaseUrl string
		Port    string
	}
}

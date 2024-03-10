package logger

type Config struct {
	Logger      string `koanf:"logger"`
	Development bool   `koanf:"development"`
	Encoding    string `koanf:"encoding"`
	Level       string `koanf:"level"`
}

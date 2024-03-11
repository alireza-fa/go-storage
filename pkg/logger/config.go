package logger

type Config struct {
	Logger      string `koanf:"logger"`
	Development bool   `koanf:"development"`
	Encoding    string `koanf:"encoding"`
	Level       string `koanf:"level"`
	Seq         struct {
		ApiKey  string `koanf:"api_key"`
		BaseUrl string `koanf:"base_url"`
		Port    string `koanf:"port"`
	} `koanf:"seq"`
}

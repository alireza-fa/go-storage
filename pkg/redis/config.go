package redis

type Config struct {
	Host               string
	Port               string
	Password           string
	Db                 int
	DialTimeout        int
	ReadTimeout        int
	WriteTimeout       int
	PoolSize           int
	PoolTimeout        int
	IdleCheckFrequency int
}

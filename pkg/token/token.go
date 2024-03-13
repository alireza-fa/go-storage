package token

import "time"

type Config struct {
	PrivatePem        string        `koanf:"private_pem"`
	PublicPem         string        `koanf:"public_pem"`
	AccessExpiration  time.Duration `koanf:"access_expiration"`
	RefreshExpiration time.Duration `koanf:"refresh_expiration"`
}

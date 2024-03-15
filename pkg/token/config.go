package token

import (
	"time"
)

type Config struct {
	PrivatePem        string
	PublicPem         string
	AccessExpiration  time.Duration
	RefreshExpiration time.Duration
}

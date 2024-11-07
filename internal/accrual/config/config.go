package config

import (
	"time"
)

type PostgresConfig struct {
	DatabaseDSN    string
	ConnectTimeout time.Duration
}

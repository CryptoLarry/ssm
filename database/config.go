package database

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config structures database configuration
type Config struct {
	Dialect string `default:"postgres"`
	Host    string `default:"127.0.0.1"`
	Port    int    `default:"5432"`

	Database string `default:"ssm"`
	Username string `default:"postgres"`
	Password string `default:"password"`

	DebugLogging bool `default:"false" split_words:"true"`
	SSLEnabled   bool `default:"false" split_words:"true"`
	Connections  int  `default:"40"`

	WipeConfirm bool `default:"false" split_words:"true"`
}

// NewConfig creates a new config by parsing environment variables
func NewConfig() *Config {
	var config Config
	envconfig.MustProcess("database", &config)
	return &config
}

// ConnectionString returns a database connection string
func (config *Config) ConnectionString() string {
	// return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.Username, config.Database, config.Password, config.SSLString())
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.Username, config.Password, config.Host, config.Port, config.Database, config.SSLString())
}

// SSLString string for SSL enable/disable
func (config *Config) SSLString() string {
	if config.SSLEnabled {
		return "require"
	}

	return "disable"
}

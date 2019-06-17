package sqlstore

// Config is the configuration for the sql storage client
type Config struct {
	DriverNameHost  string
	Port            string
	DBName          string
	Username        string
	Password        string
	RecordTableName string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		DriverName: "",
	}
}

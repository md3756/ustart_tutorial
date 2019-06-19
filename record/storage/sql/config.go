package sqlstore

// Config is the configuration for the sql storage client
type Config struct {
	DriverName      string
	Port            string
	Host            string
	DBName          string
	Username        string
	Password        string
	RecordTableName string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		DriverName:      "",
		Port:            "",
		Host:            "",
		DBName:          "",
		Username:        "",
		Password:        "",
		RecordTableName: "",
	}
}

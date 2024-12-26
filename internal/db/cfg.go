package db

// holds database configuration
type Config struct {
	Sqlite *ConfigSqlite `json:"sqlite" validate:"required"`
}

package database

type DatabaseConfig struct {
	User         string `env:"DB_USER" env-default:"postgres"`
	Password     string `env:"DB_PASSWORD" env-default:"postgres"`
	Host         string `env:"DB_HOST" env-default:"localhost"`
	Name         string `env:"DB_NAME" env-default:"postgres"`
	Port         string `env:"DB_PORT" env-default:"5432"`
	Options      string `env:"DB_OPTIONS" env-default:""`
	MigrationDir string `env:"DB_MIGRATION_DIR" env-default:"./migrations"`
}

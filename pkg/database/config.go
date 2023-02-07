package database

type DatabaseConfig struct {
	User         string `envconfig:"DB_USER" env-default:"postgres"`
	Password     string `envconfig:"DB_PASSWORD" env-default:"postgres"`
	Host         string `envconfig:"DB_HOST" env-default:"localhost"`
	Name         string `envconfig:"DB_NAME" env-default:"postgres"`
	Port         string `envconfig:"DB_PORT" env-default:"5432"`
	Options      string `envconfig:"DB_OPTIONS" env-default:""`
	MigrationDir string `envconfig:"DB_MIGRATION_DIR" env-default:"./migrations"`
}

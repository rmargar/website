package repository

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/rmargar/website/pkg/database"
	"github.com/rmargar/website/pkg/domain"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gotest.tools/assert"
)

const (
	postgresUser     = "rmargar"
	postgresPassword = "rmargar"
	postgresDB       = "rmargar"
	postgresVersion  = "14.1"
	twoMinutes       = 2 * 60
)

type DatabaseTestSuite struct {
	suite.Suite
	pool     *dockertest.Pool
	resource *dockertest.Resource
	db       *gorm.DB
}

func (s *DatabaseTestSuite) SetupSuite() {
	db, host := s.SetupPostgres()

	database.MigrateUp(db, GetConfig(host, s.resource.GetPort("5432/tcp")))
}

func (s *DatabaseTestSuite) SetupPostgres() (*sql.DB, string) {
	s.pool = NewPool()
	s.resource = RunPostgres(s.pool)
	if err := s.resource.Expire(twoMinutes); err != nil {
		panic(fmt.Errorf("[resource leaking] failed to set container expire: %w", err))
	}
	host := GetHost(s.pool)
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&client_encoding=UTF8",
		postgresUser, postgresPassword, host, s.resource.GetPort("5432/tcp"), postgresDB,
	)
	if err := s.pool.Retry(HealthCheck("postgres", dsn)); err != nil {
		panic(fmt.Errorf("could not connect to postgres (%s): %w", dsn, err))
	}
	s.db = GetGormDB(GetConfig(host, s.resource.GetPort("5432/tcp")))
	db, err := s.db.DB()
	if err != nil {
		panic(err)
	}
	return db, host
}

func (s *DatabaseTestSuite) TearDownSuite() {
	if err := s.pool.Purge(s.resource); err != nil {
		log.Errorf("could not purge resource: %s", err)
	}
}

func (s *DatabaseTestSuite) SetupTest() {
}

func (s *DatabaseTestSuite) TearDownTest() {
	RemoveAllData(s.db)
}

func GetConfig(host string, port string) *database.DatabaseConfig {
	return &database.DatabaseConfig{
		User:         postgresUser,
		Password:     postgresPassword,
		Host:         host,
		Name:         postgresDB,
		Port:         port,
		Options:      "sslmode=disable",
		MigrationDir: "../../migrations",
	}
}

func HealthCheck(driverName string, dsn string) func() error {
	return func() error {
		open, err := sql.Open(driverName, dsn)
		if err != nil {
			return fmt.Errorf("failed connect to DB: %w", err)
		}
		if err = open.Ping(); err != nil {
			return fmt.Errorf("failed on ping: %w", err)
		}
		return nil
	}
}

func NewPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(fmt.Errorf("could not connect to docker: %w", err))
	}
	return pool
}

func GetHost(pool *dockertest.Pool) string {
	u, err := url.Parse(pool.Client.Endpoint())
	if err != nil {
		panic(fmt.Errorf("invalid endpoint: %w", err))
	}
	// we don't need port, so that is why we don't check error
	host, _, _ := net.SplitHostPort(u.Host) // nolint:errcheck
	if host == "" {
		host = "localhost"
	}
	return host
}

func RunPostgres(pool *dockertest.Pool) *dockertest.Resource {
	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        postgresVersion,
			Env: []string{
				"POSTGRES_USER=" + postgresUser,
				"POSTGRES_PASSWORD=" + postgresPassword,
				"POSTGRES_DB=" + postgresDB,
			},
		}, func(hc *docker.HostConfig) {
			hc.AutoRemove = true
			hc.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to run postgres: %w", err))
	}
	return resource
}

func GetGormDB(cfg *database.DatabaseConfig) *gorm.DB {
	newLogger := logger.New(
		log.New(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s %s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
		cfg.Options,
	)
	db, err := gorm.Open(
		postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		},
	)
	if err != nil {
		panic(err)
	}
	return db
}

func RemoveAllData(db *gorm.DB) {
	tx := db.Begin()
	tx.Exec("DELETE FROM posts")
	tx.Commit()
}

func (s *DatabaseTestSuite) TestPostRepoSql_New() {

	type args struct {
		record domain.Post
	}

	tests := []struct {
		name    string
		args    args
		want    domain.Post
		wantErr bool
	}{
		{
			name:    "Should insert one without ID",
			args:    args{record: domain.Post{Author: "Test", Title: "Test", Content: "Test", URLPath: "Test1"}},
			want:    domain.Post{ID: 1, Author: "Test", Title: "Test", Content: "Test", URLPath: "Test"},
			wantErr: false,
		},
		{
			name:    "Should insert one with ID",
			args:    args{record: domain.Post{ID: 4, Author: "Test", Title: "Test", Content: "Test", URLPath: "Test2"}},
			want:    domain.Post{ID: 4, Author: "Test", Title: "Test", Content: "Test", URLPath: "Test2"},
			wantErr: false,
		},
		{
			name:    "Should return primary key error",
			args:    args{record: domain.Post{ID: 1, Author: "Test", Title: "Test", Content: "Test", URLPath: "Test3"}},
			want:    domain.Post{ID: 1, Author: "Test", Title: "Test", Content: "Test", URLPath: "Test3"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		repo := NewPostRepository(s.db)
		got, err := repo.New(tt.args.record)
		if err != nil && !tt.wantErr {
			s.T().Errorf("PostRepoSql.New() error = %v, wantErr %v", err, nil)
			return
		}

		assert.Equal(s.T(), got.Author, tt.want.Author)
		assert.Equal(s.T(), got.Title, tt.want.Title)
		assert.Equal(s.T(), got.ID, tt.want.ID)
	}
}

func (s *DatabaseTestSuite) TestPostRepoSql_GetAll() {

	tests := []struct {
		name         string
		wantErr      bool
		addedEntries int
	}{
		{
			name:         "Should retrieve all",
			wantErr:      false,
			addedEntries: 2,
		},
		{
			name:         "Should retrieve none",
			wantErr:      false,
			addedEntries: 0,
		},
	}

	for _, tt := range tests {
		want := []domain.Post{}
		tx := s.db.Begin()
		for i := 1; i <= tt.addedEntries; i++ {
			stm := fmt.Sprintf(`INSERT INTO posts (id, created_at, updated_at, author, title, content, tags, url_path) VALUES (%d, '2023-01-31 21:06:22.329000 +00:00', '2023-01-31 21:06:24.213000 +00:00', 'Test_%d', 'Test_%d', 'Test_%d', '{}', 'test%d');`, i, i, i, i, i)
			tx.Exec(stm)
			want = append(
				want,
				domain.Post{
					ID:      i,
					Author:  fmt.Sprintf("Test_%d", i),
					Content: fmt.Sprintf("Test_%d", i),
					Title:   fmt.Sprintf("Test_%d", i),
				},
			)
		}
		tx.Commit()
		repo := NewPostRepository(s.db)
		got, err := repo.GetAll()
		if err != nil && !tt.wantErr {
			s.T().Errorf("PostRepoSql.GetAll() error = %v, wantErr %v", err, nil)
			return
		}
		assert.Equal(s.T(), len(got), len(want))
		s.TearDownTest()
	}
}

func (s *DatabaseTestSuite) TestPostRepoSql_SearchByTitle() {
	tx := s.db.Begin()
	for i := 1; i <= 2; i++ {
		stm := fmt.Sprintf(`INSERT INTO posts (id, created_at, updated_at, author, title, content, tags, url_path) VALUES (%d, '2023-01-31 21:06:22.329000 +00:00', '2023-01-31 21:06:24.213000 +00:00', 'Test_%d', 'Test_%d', 'Test_%d', '{}', 'test%d');`, i, i, i, i, i)
		tx.Exec(stm)
	}
	tx.Commit()

	type args struct {
		title string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
		want    []domain.Post
	}{
		{
			name:    "Should retrieve all",
			wantErr: false,
			args:    args{title: "Test"},
			want: []domain.Post{
				{
					ID:      1,
					Author:  "Test_1",
					Content: "Test_1",
					Title:   "Test_1",
					URLPath: "test1",
				},
				{
					ID:      2,
					Author:  "Test_2",
					Content: "Test_2",
					Title:   "Test_2",
					URLPath: "test2",
				},
			},
		},
		{
			name:    "Should retrieve none",
			wantErr: false,
			args:    args{title: "I don't exist"},
			want:    []domain.Post{},
		},
		{
			name:    "Should retrieve one",
			wantErr: false,
			args:    args{title: "Test_2"},
			want: []domain.Post{
				{
					ID:      2,
					Author:  "Test_2",
					Content: "Test_2",
					Title:   "Test_2",
				},
			},
		},
	}

	for _, tt := range tests {
		repo := NewPostRepository(s.db)
		got, err := repo.SearchByTitle(tt.args.title)
		if err != nil && !tt.wantErr {
			s.T().Errorf("PostRepoSql.SearchByTitle() error = %v, wantErr %v", err, nil)
			return
		}
		assert.Equal(s.T(), len(got), len(tt.want))
	}
}

func (s *DatabaseTestSuite) TestPostRepoSql_GetByTag() {
	tx := s.db.Begin()
	for i := 1; i <= 2; i++ {
		stm := fmt.Sprintf(`INSERT INTO posts (id, created_at, updated_at, author, title, content, tags, url_path) VALUES (%d, '2023-01-31 21:06:22.329000 +00:00', '2023-01-31 21:06:24.213000 +00:00', 'Test_%d', 'Test_%d', 'Test_%d', '{"Test"}', 'test%d');`, i, i, i, i, i)
		tx.Exec(stm)
	}
	tx.Commit()

	type args struct {
		tag string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
		want    []domain.Post
	}{
		{
			name:    "Should retrieve all",
			wantErr: false,
			args:    args{tag: "Test"},
			want: []domain.Post{
				{
					ID:      1,
					Author:  "Test_1",
					Content: "Test_1",
					Title:   "Test_1",
					URLPath: "test1",
				},
				{
					ID:      2,
					Author:  "Test_2",
					Content: "Test_2",
					Title:   "Test_2",
					URLPath: "test2",
				},
			},
		},
		{
			name:    "Should retrieve none",
			wantErr: false,
			args:    args{tag: "sometag"},
			want:    []domain.Post{},
		},
	}

	for _, tt := range tests {
		repo := NewPostRepository(s.db)
		got, err := repo.GetByTag(tt.args.tag)
		if err != nil && !tt.wantErr {
			s.T().Errorf("PostRepoSql.GetByTag() error = %v, wantErr %v", err, nil)
			return
		}
		assert.Equal(s.T(), len(got), len(tt.want))
	}
}

func TestIntegrationDatabaseTestSuite(t *testing.T) {
	suite.Run(t, &DatabaseTestSuite{})
}

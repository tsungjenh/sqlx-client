package sqlx

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Username          string        `json:"username" mapstructure:"username"`
	Password          string        `json:"password" mapstructure:"password"`
	Host              string        `json:"host" mapstructure:"host"`
	Port              string        `json:"port" mapstructure:"port"`
	Name              string        `json:"db_name" mapstructure:"db_name"`
	DbMaxIdleConns    int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	DbMaxOpenConns    int           `json:"max_open_conns" mapstructure:"max_open_conns"`
	DbConnMaxLifetime time.Duration `json:"conn_max_life_time" mapstructure:"conn_max_life_time"`
}

type databaseClient struct {
	db *sqlx.DB
}

var dbHandles = map[string]databaseClient{}

func AddDb(dbName string, dbHandle *sqlx.DB) {
	dbHandles[strings.ToLower(dbName)] = newDatabaseClient(dbHandle)
}

func newDatabaseClient(db *sqlx.DB) *databaseClient {
	return &databaseClient{
		db: db,
	}
}

func InitialDatabase(databases map[string]*DatabaseConfig) error {
	for dbName, c := range databases {
		dsn := fmt.Sprintf(
			DsnFormat,
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.DbName,
		)
		db, err := sqlx.Connect("mysql", dsn)
		if err != nil {
			return err
		}

		if c.DbMaxIdleConns == 0 {
			c.DbMaxIdleConns = defaultDbMaxIdleConns
		}
		if c.DbMaxOpenConns == 0 {
			c.DbMaxOpenConns = defaultDbMaxOpenConns
		}
		if c.DbConnMaxLifetime == 0 {
			c.DbConnMaxLifetime = defaultDbConnMaxLifetime
		}
		db.SetConnMaxLifetime(c.DbConnMaxLifetime)
		db.SetMaxIdleConns(c.DbMaxIdleConns)
		db.SetMaxOpenConns(c.DbMaxOpenConns)
		AddDb(dbName, db)
	}
	return nil
}

func GetDb(dbName string) (databaseClient, error) {
	db, ok := dbHandles[dbName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("db not exists: %s", dbName))
	}

	return db, nil
}

func CloseDbs() {
	for _, dbHandle := range dbHandles {
		_ = dbHandle.Close()
	}
}

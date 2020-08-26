# Extends github.com/jmoiron/sqlx package

## Installation

~~~sh
go get -u github.com/tsungjenh/sqlx-client
~~~

## Usage

~~~go
import github.com/tsungjenh/sqlx-client

config = map[string]DatabaseConfig{
	"db1": {
		Username: "root",
		Password: "123456",
		Host: "localhost",
		Port: "3306",
		DBName: "foo",
		DbMaxIdleConns: 10,
		DbMaxOpenConns: 100,
		DbConnMaxLifetime: 30,
	},
	"db2": {
		Username: "root",
		Password: "123456",
		Host: "localhost",
		Port: "3306",
		DBName: "bar",
		DbMaxIdleConns: 10,
		DbMaxOpenConns: 100,
		DbConnMaxLifetime: 30,
	},
	// ...
}

dbClient, err := GetDb("db1")
if err != nil {
	// db not found, handle it.	
}
// dbClient.db.Exec(sql, params...)

~~~
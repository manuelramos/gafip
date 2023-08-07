// db_factory.go
package db

import (
	"os"
)

func NewDBConnector() DBConnector {
    dbType := os.Getenv("DB_TYPE")
    switch dbType {
    case "postgres":
        return &PostgresConnector{}
	case "sqlite":
		return &SQLiteConnector{}
    default:
        return &SQLiteConnector{}
    }
}

// db_connector.go
package db

type DBConnector interface {
    Connect() error
    Close() error
    Exec(query string, args ...interface{}) error
    Query(query string, args ...interface{}) ([]map[string]interface{}, error)
    BeginTransaction() (Transaction, error)
}

type Transaction interface {
    Exec(query string, args ...interface{}) error
    Query(query string, args ...interface{}) ([]map[string]interface{}, error)
    Commit() error
    Rollback() error
}

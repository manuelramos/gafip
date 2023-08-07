// postgres_connector.go
package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresConnector struct {
    db *sql.DB
}

func (p *PostgresConnector) Connect() error {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return err
    }

    p.db = db
    return nil
}

func (p *PostgresConnector) Close() error {
    return p.db.Close()
}

func (p *PostgresConnector) Exec(query string, args ...interface{}) error {
    _, err := p.db.Exec(query, args...)
    return err
}

func (p *PostgresConnector) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
    rows, err := p.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []map[string]interface{}
    columns, _ := rows.Columns()
    count := len(columns)
    values := make([]interface{}, count)
    valuePtrs := make([]interface{}, count)

    for rows.Next() {
        for i := range columns {
            valuePtrs[i] = &values[i]
        }

        rows.Scan(valuePtrs...)
        row := make(map[string]interface{})

        for i, col := range columns {
            var v interface{}
            val := values[i]

            b, ok := val.([]byte)
            if ok {
                v = string(b)
            } else {
                v = val
            }

            row[col] = v
        }

        results = append(results, row)
    }

    return results, nil
}

func (p *PostgresConnector) BeginTransaction() (Transaction, error) {
    tx, err := p.db.Begin()
    if err != nil {
        return nil, err
    }
    return &PostgresTransaction{tx}, nil
}

type PostgresTransaction struct {
    tx *sql.Tx
}

func (pt *PostgresTransaction) Exec(query string, args ...interface{}) error {
    _, err := pt.tx.Exec(query, args...)
    return err
}

func (pt *PostgresTransaction) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
    rows, err := pt.tx.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []map[string]interface{}
    columns, _ := rows.Columns()
    count := len(columns)
    values := make([]interface{}, count)
    valuePtrs := make([]interface{}, count)

    for rows.Next() {
        for i := range columns {
            valuePtrs[i] = &values[i]
        }

        rows.Scan(valuePtrs...)
        row := make(map[string]interface{})

        for i, col := range columns {
            var v interface{}
            val := values[i]

            b, ok := val.([]byte)
            if ok {
                v = string(b)
            } else {
                v = val
            }

            row[col] = v
        }

        results = append(results, row)
    }

    return results, nil
}

func (pt *PostgresTransaction) Commit() error {
    return pt.tx.Commit()
}

func (pt *PostgresTransaction) Rollback() error {
    return pt.tx.Rollback()
}

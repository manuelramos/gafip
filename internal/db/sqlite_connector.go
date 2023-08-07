// sqlite_connector.go
package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteConnector struct {
	db *sql.DB
}

func NewSQLiteConnector() DBConnector {
	return &SQLiteConnector{}
}

func (s *SQLiteConnector) Connect() error {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		// If no path is provided, we use a relative path to the current directory
		dbPath = filepath.Join(".", "db.sqlite")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *SQLiteConnector) Close() error {
	return s.db.Close()
}

func (s *SQLiteConnector) Exec(query string, args ...interface{}) error {
	_, err := s.db.Exec(query, args...)
	return err
}

func (s *SQLiteConnector) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := s.db.Query(query, args...)
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

func (s *SQLiteConnector) BeginTransaction() (Transaction, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	return &SQLiteTransaction{tx}, nil
}

type SQLiteTransaction struct {
	tx *sql.Tx
}

func (st *SQLiteTransaction) Exec(query string, args ...interface{}) error {
	_, err := st.tx.Exec(query, args...)
	return err
}

func (st *SQLiteTransaction) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := st.tx.Query(query, args...)
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

func (st *SQLiteTransaction) Commit() error {
	return st.tx.Commit()
}

func (st *SQLiteTransaction) Rollback() error {
	return st.tx.Rollback()
}

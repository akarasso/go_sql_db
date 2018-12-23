package ft_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConn struct {
	conn	*sql.DB
}

func Connect(dsn string) (*DatabaseConn, error) {
	var err error
	db := new(DatabaseConn)
    db.conn, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
    }
    if err = db.conn.Ping(); err != nil {
		return nil, err
    }
	return db, nil
}

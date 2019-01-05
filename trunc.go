package ft_database

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func (db DatabaseConn)Trunc(table string) (sql.Result, error) {
	req := "TRUNCATE `"+table+"`;"
	stmt, err := db.Conn.Prepare(req)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec()
}

package ft_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func (db DatabaseConn) Delete(table string, data map[string]interface{}) (sql.Result, error) {
    var req string
    var dataValue []interface{}

    req = "DELETE FROM `" + table + "`"

    if err := addFieldsClause(&req, data, &dataValue); err != nil {
        return nil, err
    }
    stmt, err := db.Conn.Prepare(req)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    return stmt.Exec(dataValue...)
}

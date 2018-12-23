package ft_database

import (
    "strings"
    "errors"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func addInsertValue(req *string, data map[string]interface{}, dataValue *[]interface{}) error {
    var cols []string
    var vals []string

    if values, exist := data["values"]; exist {
        switch values.(type) {
        case map[string]interface{}:
            for key, val := range values.(map[string]interface{}) {
                cols = append(cols, "`" + key + "`")
                vals = append(vals, "?")
                (*dataValue) = append((*dataValue), val)
            }
            (*req) += strings.Join(cols, ", ") + ") VALUES (" + strings.Join(vals, ", ") + ")"
            return nil
        default:
            return errors.New("'where' clause only accept map[string]interface{}")
        }
    }
    return errors.New("'values' unset")
}

func (db DatabaseConn)Insert(table string, data map[string]interface{}) (sql.Result, error) {
    var req string
    var dataValue []interface{}

    req = "INSERT INTO `" + table + "` ("
    if err := addInsertValue(&req, data, &dataValue); err != nil {
        return nil, err
    }
    stmt, err := db.conn.Prepare(req)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    return stmt.Exec(dataValue...)
}

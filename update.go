package ft_database
import (
    "strings"
    "errors"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func addUpdateFields(req *string, data map[string]interface{}, dataValue *[]interface{}) (error) {
    if fields, exist := data["fields"]; exist {
        switch fields.(type) {
        case map[string]interface{}:
            var dataCond []string
            for key, val := range fields.(map[string]interface{}) {
                dataCond = append(dataCond, "`" + key + "` = ?")
                (*dataValue) = append((*dataValue), val)
            }
            (*req) += strings.Join(dataCond, ", ")
            return nil
        default:
            return errors.New("'fields' clause only accept map[string]interface{}")
        }
    }
    return errors.New("'fields' is unset")
}

func (db DatabaseConn)Update(table string, data map[string]interface{}) (sql.Result, error) {
    var req string
    var dataValue []interface{}

    req = "UPDATE `" + table + "` SET "
    if err := addUpdateFields(&req, data, &dataValue); err != nil {
        return nil, err
    }
    if err := addFieldsClause(&req, data, &dataValue); err != nil {
        return nil, err
    }
    stmt, err := db.conn.Prepare(req)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    return stmt.Exec(dataValue...)
}

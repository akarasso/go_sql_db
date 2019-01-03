package ft_database

import (
    "strings"
    "errors"
)

func addFieldsSelect(req *string, data map[string]interface{}) (error) {
    if fields, exist := data["fields"]; exist {
        switch fields.(type) {
        case string:
            (*req) += fields.(string)
        case []string:
            (*req) += strings.Join(fields.([]string), ", ")
        default:
            return errors.New("'fields' clause only accept string/Array")
        }
        return nil
    } else {
        (*req) += "*"
    }
    return nil
}

func addFieldsClause(req *string, data map[string]interface{}, dataValue *[]interface{}) (error) {
    if where, exist := data["where"]; exist {
        switch where.(type) {
        case string:
            (*req) += " WHERE " +  where.(string)
            if whereval, ex := data["values"]; ex {
                switch whereval.(type) {
                case []interface{}:
                    for _, v := range whereval.([]interface{}) {
                        (*dataValue) = append((*dataValue), v)
                    }
                default:
                    return errors.New("'where' clause string must accept []interface values")
                }
            } else {
                return errors.New("'whereval' must be in values")
            }
            return nil
        case map[string]interface{}:
            var dataCond []string
            for key, val := range where.(map[string]interface{}) {
                dataCond = append(dataCond, "`" + key + "` = ?")
                (*dataValue) = append((*dataValue), val)
            }
            (*req) += " WHERE " +  strings.Join(dataCond, " AND ")
            return nil
        default:
            return errors.New("'where' clause only accept map[string]interface{}")
        }
    }
    return errors.New("'where' is unset")
}

func (db DatabaseConn)SelectFirst(table string, data map[string]interface{}) (map[string]interface{}, error) {
    var req string
    var dataValue []interface{}

    req = "SELECT "
    if err := addFieldsSelect(&req, data); err != nil {
        return nil, err
    }
    req += " FROM `" + table + "`"
    if err := addFieldsClause(&req, data, &dataValue); err != nil {
        return nil, err
    }
    stmt, err := db.Conn.Prepare(req)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    rows, err := stmt.Query(dataValue...)
    if err != nil {
        return nil, err
    }
    return Read_first_result(rows)
}

func (db DatabaseConn)Select(table string, data map[string]interface{}) ([]map[string]interface{}, error) {
    var req string
    var dataValue []interface{}

    req = "SELECT "
    if err := addFieldsSelect(&req, data); err != nil {
        return nil, err
    }
    req += " FROM `" + table + "`"
    if err := addFieldsClause(&req, data, &dataValue); err != nil {
        return nil, err
    }
    stmt, err := db.Conn.Prepare(req)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    rows, err := stmt.Query(dataValue...)
    if err != nil {
        return nil, err
    }
    return Read_result(rows)
}

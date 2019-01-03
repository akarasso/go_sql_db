package ft_database

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Read_first_result(rows *sql.Rows) (map[string]interface{}, error) {
    cols, _ := rows.Columns()
    for rows.Next() {
        columns := make([]interface{}, len(cols))
        columnPointers := make([]interface{}, len(cols))
        for i, _ := range columns {
            columnPointers[i] = &columns[i]
        }
        if err := rows.Scan(columnPointers...); err != nil {
            return nil, err
        }
        m := make(map[string]interface{})
        for i, colName := range cols {
            val := columnPointers[i].(*interface{})
            switch (*val).(type) {
			case []byte:
				m[colName] = string((*val).([]byte))
			default:
				m[colName] = *val
			}
        }
        return m, nil
    }
    return nil, nil
}

func Read_result(rows *sql.Rows) ([]map[string]interface{}, error) {
    var res []map[string]interface{}
    cols, _ := rows.Columns()
    for rows.Next() {
        columns := make([]interface{}, len(cols))
        columnPointers := make([]interface{}, len(cols))
        for i, _ := range columns {
            columnPointers[i] = &columns[i]
        }
        if err := rows.Scan(columnPointers...); err != nil {
            return nil, err
        }
        m := make(map[string]interface{})
        for i, colName := range cols {
            val := columnPointers[i].(*interface{})
            switch (*val).(type) {
			case []byte:
				m[colName] = string((*val).([]byte))
			default:
				m[colName] = *val
			}
        }
        res = append(res, m)
    }
    return res, nil
}

package ft_database

func (db DatabaseConn)GetColumns(table string) ([]string, error) {
    var req string

    req = "SELECT column_name FROM information_schema.columns WHERE table_name = ?"
	stmt, err := db.Conn.Prepare(req)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    rows, err := stmt.Query(table)
    if err != nil {
        return nil, err
    }
	var cols []string
	for rows.Next() {
		var colName string
		if err := rows.Scan(&colName); err != nil {
            return nil, err
        }
		cols = append(cols, colName)
	}
	return cols, nil
}
// SELECT column_name FROM information_schema.columns WHERE table_name = 'interest'

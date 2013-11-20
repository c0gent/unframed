package unframed

import (
	"database/sql"
	"github.com/nsan1129/auctionLog/log"
)

//"github.com/nsan1129/auctionLog/log"

type DataAdapter struct {
	SessionManager
}

func (a DataAdapter) Query(stmt *UfStmt, p ...interface{}) *sql.Rows {
	rows, err := stmt.Query(p...)
	if err != nil {
		panic(err)
	}
	return rows
}

func (a DataAdapter) QueryRow(stmt *UfStmt, p ...interface{}) *sql.Row {
	row := stmt.QueryRow(p...)
	return row
}

func (a DataAdapter) ScanRows(rows *sql.Rows, f ...interface{}) {
	err := rows.Scan(f...)
	if err != nil {
		log.Error(err)
	}
}

func (a DataAdapter) ScanRow(row *sql.Row, f ...interface{}) {
	err := row.Scan(f...)
	if err != nil {
		log.Error(err)
	}
}

func (a DataAdapter) Exec(stmt *UfStmt, p ...interface{}) {
	_, err := stmt.Exec(p...)
	if err != nil {
		log.Error(err)
	}
}

func (a DataAdapter) CrazyQueryScan(scanner func(rows *sql.Rows), stmt *UfStmt, p ...interface{}) {
	rs := a.Query(stmt, 20)
	defer rs.Close()

	for rs.Next() {
		scanner(rs)
	}
}

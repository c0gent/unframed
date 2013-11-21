package unframed

import (
	"database/sql"
	"github.com/nsan1129/auctionLog/log"
	"reflect"
)

//"github.com/nsan1129/auctionLog/log"

type DataAdapter struct {
	SessionManager
}

func (a DataAdapter) query(stmt *UfStmt, p ...interface{}) *sql.Rows {
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

func (a DataAdapter) scan(rows *sql.Rows, f ...interface{}) {
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

func (a DataAdapter) Query(newStruct func() interface{}, stmt *UfStmt, p ...interface{}) {
	rs := a.query(stmt, p...)
	defer rs.Close()

	for rs.Next() {
		a.scan(rs, getTargets(newStruct())...)
	}
}

type structField struct {
	name  string
	index int
}

type structFields struct {
	fields map[string]*structField
}

/*

func getFields(ds interface{}) (sfs *structFields) {
	dsType := reflect.TypeOf(ds)

	if dsType.Kind() == reflect.Ptr {
		dsType = dsType.Elem()
	}

	if dsType.Kind() != reflect.Struct {
		log.Error(dsType.Kind(), "isn't a struct.")
		return
	}

	sfs = new(structFields)
	sfs.fields = make(map[string]*structField)

	for i := 0; i < dsType.NumField(); i++ {

		sf := dsType.Field(i)

		if sf.PkgPath != "" {
			continue
		}

		sfs.fields[sf.Name] = &structField{
			name:  sf.Name,
			index: i,
		}
	}

	return
}

*/

func getTargets(ds interface{}) (stf []interface{}) {

	dsType := reflect.TypeOf(ds)
	dsValue := reflect.ValueOf(ds).Elem()

	if dsType.Kind() == reflect.Ptr {
		dsType = dsType.Elem()
	}
	/*
		if dsType.Kind() != reflect.Struct {
			log.Error(dsType.Kind(), "isn't a struct.")
			return
		}
	*/

	for i := 0; i < dsType.NumField(); i++ {

		structField := dsType.Field(i)

		if structField.PkgPath != "" {
			continue
		}

		fieldAddr := dsValue.Field(i).Addr().Interface()
		stf = append(stf, fieldAddr)
	}
	return
}

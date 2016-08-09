package unframed

import (
	"database/sql"
	"errors"
	"github.com/c0gent/unframed/log"
	"reflect"
)

//"github.com/nsan1129/auctionLog/log"

type DataAdapter struct {
	Net *NetHandle
}

func (a *DataAdapter) queryStmt(stmt *UfStmt, p ...interface{}) *sql.Rows {
	rows, err := stmt.Query(p...)
	if err != nil {
		panic(err)
	}
	return rows
}

func (a *DataAdapter) QueryRow(stmt *UfStmt, p ...interface{}) *sql.Row {
	row := stmt.QueryRow(p...)
	return row
}

func (a *DataAdapter) scan(rows *sql.Rows, f ...interface{}) {
	err := rows.Scan(f...)
	if err != nil {
		log.Error(err)
	}
}

func (a *DataAdapter) ScanRow(row *sql.Row, f ...interface{}) {
	err := row.Scan(f...)
	if err != nil {
		log.Error(err)
	}
}

func (a *DataAdapter) Insert(stmt *UfStmt, p ...interface{}) (id int) {
	err := stmt.QueryRow(p...).Scan(&id)
	if err != nil {
		log.Error(err)
	}
	return id
}

func (a *DataAdapter) Exec(stmt *UfStmt, p ...interface{}) {
	_, err := stmt.Exec(p...)
	if err != nil {
		log.Error(err)
	}
	
}

func (a *DataAdapter) Query(newStruct func() interface{}, stmt *UfStmt, p ...interface{}) {
	if stmt == nil {
		log.Error(errors.New("unframed.DataAdapter.Query called with nil *UfStmt"))
		return
	}
	rs := a.queryStmt(stmt, p...)
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

func getTargets(ds interface{}) (stf []interface{}) {

	/*
		dsType := reflect.TypeOf(ds)
		if dsType.Kind() == reflect.Ptr {
			dsType = dsType.Elem()
		}

		if dsType.Kind() != reflect.Struct {
			log.Error(dsType.Kind(), "isn't a struct.")
			return
		}
	*/

	dsValue := reflect.ValueOf(ds).Elem()

	for i := 0; i < dsValue.NumField(); i++ {

		/*
			structField := dsType.Field(i)

			if structField.PkgPath != "" {
				continue
			}
		*/

		stf = append(stf, dsValue.Field(i).Addr().Interface())
	}
	return
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

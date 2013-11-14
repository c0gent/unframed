package unframed

import (
	"database/sql"
	"github.com/gorilla/sessions"
	//"encoding/gob"
	"github.com/gorilla/schema"
)

type dbDialogue int

const (
	DbdPg    dbDialogue = 1001
	dbdPg    dbDialogue = 1001
	DbdMysql dbDialogue = 1002
	DbdMondo dbDialogue = 1003
	DbdSqlit dbDialogue = 1004
	DbdNosql dbDialogue = 1005
	DbdGaeds dbDialogue = 1006
	DbdOther dbDialogue = 1007
)

var Dbd struct {
	Pg,
	Mysql,
	Mondo,
	Sqlit,
	Nosql,
	Gaeds,
	Other dbDialogue
}

var cookieName string = "halloween"
var cookieStore *sessions.CookieStore = sessions.NewCookieStore([]byte("candy"))

// --------- DATABASE HANDLER ----------

type DB struct {
	*sql.DB
	cdd dbDialogue
	*StatementStore
	Decoder *schema.Decoder
}

func (d *DB) init(driverName string, dataSourceName string) *DB {
	var err error
	d.DB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	d.initDbd()
	d.cdd = Dbd.Pg

	d.StatementStore = new(StatementStore).init(d)

	return d
}

func (db *DB) initDbd() {
	Dbd.Pg = 1001
	Dbd.Mysql = 1002
	Dbd.Mondo = 1003
	Dbd.Sqlit = 1004
	Dbd.Nosql = 1005
	Dbd.Gaeds = 1006
	Dbd.Other = 1007
}

func NewDB(driverName string, dataSourceName string) *DB {

	nd := new(DB).init(driverName, dataSourceName)
	nd.Decoder = schema.NewDecoder()
	return nd

}

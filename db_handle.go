package unframed

import (
	"database/sql"
)

type dbDialogue int

var Dbd struct {
	Pg,
	Mysql,
	Mondo,
	Sqlit,
	Nosql,
	Gaeds,
	Other dbDialogue
}

// --------- DATABASE HANDLER ----------

type DbHandle struct {
	*sql.DB
	cdd dbDialogue
	*StatementStore
	*SessionManager
}

func (d *DbHandle) init(driverName string, dataSourceName string) *DbHandle {
	var err error
	d.DB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	d.initDbd()
	d.cdd = Dbd.Pg

	d.StatementStore = new(StatementStore).init(d)
	//d.SessionManager = new(SessionManager)

	return d
}

func (db *DbHandle) initDbd() {
	Dbd.Pg = 1001
	Dbd.Mysql = 1002
	Dbd.Mondo = 1003
	Dbd.Sqlit = 1004
	Dbd.Nosql = 1005
	Dbd.Gaeds = 1006
	Dbd.Other = 1007
}

func NewDB(driverName string, dataSourceName string) *DbHandle {

	nd := new(DbHandle).init(driverName, dataSourceName)
	return nd

}

/*
--TERMINOLOGY--
- Actions relating to Data Objects (Records). Must be Verbs. -
Go			Http		SQL					Purpose
------------------------------------------------------------
list		GET 		SELECT				display multiple records
show		GET 		SELECT				display record

create		POST		INSERT				store new record
update		POST		UPDATE				modify existing record
save		POST		INSERT/UPDATE		create/update depending on Id

compose		GET			(none)				display composition controls/tools
edit		GET			(none)				display editing controls/tools
form		GET			(none)				edit/compose depending on Id

delete		POST		DELETE				destroy existing record



- Other -
find
NewXXX()		Return a new instance of something. Customary GO shorthand for GetNewXXX().

-Other Terms-
List(noun) = Table of Data, rows(multiple), etc.

*/

package unframed

import (
	"database/sql"
	//"github.com/nsan1129/auctionLog/log"
)

// Map of [Database Dialogue Types (ints basically)] to: SQL Statement Strings
// used to hold several statement strings for use depending on the
// database dialogue being used.
type StmtStringsByDia map[dbDialogue]string

type UfStmt struct {
	*sql.Stmt
	StringsByDia StmtStringsByDia
}

type UfStmtsByName map[string]*UfStmt

type StatementStore struct {
	Stmts    UfStmtsByName
	dbHandle *DbHandle
}

func (ss *StatementStore) init(db *DbHandle) *StatementStore {
	ss.dbHandle = db
	ss.Stmts = make(UfStmtsByName)
	return ss
}

func (ss *StatementStore) PrepareStatements() {
	for key, _ := range ss.Stmts {
		var err error
		ss.Stmts[key].Stmt, err = ss.dbHandle.Prepare(ss.Stmts[key].StringsByDia[ss.dbHandle.cdd])

		if err != nil {
			panic(err)
		}
		//log.Message("PrepareStatements(): key=", key, "statement=", ss.Stmts[key].StringsByDia[ss.dbHandle.cdd])
	}
}

func (ss *StatementStore) AddStatement(name string, dbd dbDialogue, text string) {
	if ss.Stmts[name] == nil {
		ss.Stmts[name] = &UfStmt{nil, StmtStringsByDia{dbd: text}}
		//log.Message("StatementStore.AddStatement(); ss.Stmts[name]==nil", "name:", name)
	} else {
		ss.Stmts[name].StringsByDia[dbd] = text
		//log.Message("StatementStore.AddStatement(); ss.Stmts[name]!=nil", "name:", name)
	}
}

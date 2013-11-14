package unframed

import (
	"github.com/gorilla/sessions"
	"github.com/nsan1129/auctionLog/log"
	"net/http"
)

type DataModel struct {
	session      *sessions.Session
	LoginFailure bool
}

func (dm *DataModel) SetSession(r *http.Request) {
	var err error
	dm.session, err = cookieStore.Get(r, cookieName)
	if err != nil {
		log.Error(err)
	}
}

func (dm *DataModel) IsLoggedIn() bool {

	if dm.session == nil {
		return true
	}

	return !dm.session.IsNew
}

func (dm *DataModel) GetSessionValues() interface{} {
	return dm.session.Values[0]
}

func (dm *DataModel) SetSessionValues(w http.ResponseWriter, r *http.Request, in interface{}) {
	dm.session.Values[0] = in
	dm.session.Save(r, w)
}

func (dm *DataModel) InitSession(r *http.Request) *DataModel {
	dm.SetSession(r)
	return dm
}

func (dm *DataModel) SetLoginFailure(r *http.Request, fail bool) *DataModel {
	dm.LoginFailure = fail
	dm.SetSession(r)
	return dm
}

func (dm *DataModel) DeleteSession(r *http.Request, w http.ResponseWriter) {
	dm.SetSession(r)
	dm.session.Options.MaxAge = -1
	dm.session.Save(r, w)
}

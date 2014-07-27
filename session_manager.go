package unframed

import (
	"github.com/gorilla/sessions"
	"github.com/nsan1129/unframed/log"
	"net/http"
)

var cookieName string = "unframed"
var cookieStore *sessions.CookieStore = sessions.NewCookieStore([]byte("candy"))

type SessionManager struct {
	session      *sessions.Session
	LoginFailure bool
}

func (dm *SessionManager) SetSession(r *http.Request) {
	var err error
	dm.session, err = cookieStore.Get(r, cookieName)
	if err != nil {
		log.Error(err)
	}
}

func (dm *SessionManager) IsLoggedIn() bool {

	if dm.session == nil {
		return true
	}

	return !dm.session.IsNew
}

func (dm *SessionManager) GetSessionValues() interface{} {
	return dm.session.Values[0]
}

func (dm *SessionManager) SetSessionValues(w http.ResponseWriter, r *http.Request, in interface{}) {
	dm.session.Values[0] = in
	dm.session.Save(r, w)
}

func (dm *SessionManager) InitSession(r *http.Request) *SessionManager {
	dm.SetSession(r)
	return dm
}

func (dm *SessionManager) SetLoginFailure(r *http.Request, fail bool) *SessionManager {
	dm.LoginFailure = fail
	dm.SetSession(r)
	return dm
}

func (dm *SessionManager) DeleteSession(r *http.Request, w http.ResponseWriter) {
	dm.SetSession(r)
	dm.session.Options.MaxAge = -1
	dm.session.Save(r, w)
}

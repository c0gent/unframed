package unframed

import (
	"github.com/gorilla/sessions"
	"github.com/c0gent/unframed/log"
	"net/http"
)

var cookieName string = "unframed"
var cookieStore *sessions.CookieStore = sessions.NewCookieStore([]byte("candy"))

type SessionManager struct {
	Session      *sessions.Session
	LoginFailure bool
}

func NewSessionManager() *SessionManager {
	return new(SessionManager)
}

func (dm *SessionManager) SetSession(r *http.Request) {
	var err error
	var test *sessions.Session
	test, err = cookieStore.Get(r, cookieName)
	if err != nil {
		log.Error(err)
	}
	dm.Session = test
}

func (dm *SessionManager) IsLoggedIn() bool {

	if dm.Session == nil {
		return true
	}

	return !dm.Session.IsNew
}

func (dm *SessionManager) GetSessionValues() interface{} {
	return dm.Session.Values[0]
}

func (dm *SessionManager) SetSessionValues(w http.ResponseWriter, r *http.Request, in interface{}) {
	dm.Session.Values[0] = in
	dm.Session.Save(r, w)
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
	dm.Session.Options.MaxAge = -1
	dm.Session.Save(r, w)
}

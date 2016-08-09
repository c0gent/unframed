package unframed

import (
	"encoding/gob"
	"github.com/gorilla/schema"
	"github.com/c0gent/unframed/log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var DefaultPageTitle string = "Unframed Default Page Title"

type TmplDataWrapper struct {
	PageTitle string
	Data      interface{}
	Net       *NetHandle
	Sdata     []interface{}
}

type NetHandle struct {
	Decoder *schema.Decoder
	*Router
	*SessionManager
	*TemplateStore
}

func (n *NetHandle) RegType(t interface{}) {
	gob.Register(t)
}
func (n *NetHandle) ExeTmpl(w http.ResponseWriter, templateName string, templateData ...interface{}) {
	var tw *TmplDataWrapper
	tw = &TmplDataWrapper{Data: templateData, Net: n, PageTitle: DefaultPageTitle}

	/*
		for _, data := range templateData[1:] {
			log.Message(data)
		}
	*/

	/*
		if len(templateData) > 1 {
			tw.Data = templateData[0:]
			for i, data := range templateData {
				log.Message(i)
				log.Message(data)
			}
		}
	*/
	err := n.templates.ExecuteTemplate(w, templateName, tw)
	if err != nil {
		log.Error(err)
	}
}
func (n *NetHandle) PrettyDate(t time.Time, style int) (datestring string) {
	switch style {
	case 0:
		datestring = t.Format("2006/01/02 15:04:05 MST")
	case 1:
		datestring = t.Format("Jan 02 15:04")
	case 2:
		datestring = t.Format("2006-Jan-02 15:04:05 MST")
	default:
		datestring = t.Format("Monday, January 2nd, 2006 15:04:05 -0700 MST")
	}

	return
}
func (n *NetHandle) DecodeForm(target interface{}, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}
	err = n.Decoder.Decode(target, r.PostForm)
	if err != nil {
		log.Error(err)
	}
}
func (n *NetHandle) Compare(a int, b int) (x bool) {
	x = a == b
	return
}
func (n *NetHandle) NewTmplDataWrapper() (newTdw *TmplDataWrapper) {
	newTdw = new(TmplDataWrapper)
	newTdw.PageTitle = DefaultPageTitle
	return
}

func NewNet() (nn *NetHandle) {
	nn = new(NetHandle)
	nn.Decoder = schema.NewDecoder()
	nn.Router = NewRouter()
	nn.TemplateStore = NewTemplateStore()
	nn.SessionManager = NewSessionManager()

	nn.RegType(new(time.Time))

	return
}

func (n *NetHandle) TimeSince(t time.Time) string {
	si := time.Since(t).Minutes()
	return strconv.FormatFloat(si, 'f', 0, 64)
}

func (n *NetHandle) IntInStr(str string, v int, sep string) bool {
	tmp := strings.Split(str, sep)
	for _, ele := range tmp {
		if Atoi(ele) == v {
			//log.Message(ele, " = ", v)
			return true
		}
	}
	//log.Message("Nothing Found for IntInStr")
	return false
}

func (n *NetHandle) StrAppendInt(old string, val int, sep string) (new string) {
	tmp := strings.Split(old, sep)
	tmp = append(tmp, Itoa(val))
	new = strings.Join(tmp, sep)
	return
}

package unframed

import (
	"encoding/gob"
	"github.com/gorilla/schema"
	"github.com/nsan1129/unframed/log"
	"net/http"
	"strconv"
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
	tw = &TmplDataWrapper{Data: templateData[0], Net: n, PageTitle: DefaultPageTitle}
	if len(templateData) > 1 {
		tw.Sdata = templateData[1:]
	}
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

	nn.RegType(new(time.Time))

	return
}

func (n *NetHandle) TimeSince(t time.Time) (string) {
	si := time.Since(t).Minutes()
	return strconv.FormatFloat(si, 'f', 0, 64)
}
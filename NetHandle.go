package unframed

import (
	"encoding/gob"
	"github.com/gorilla/schema"
	"github.com/nsan1129/auctionLog/log"
	"net/http"
	"time"
)

type NetHandle struct {
	Decoder *schema.Decoder
	*Router
	*SessionManager
	*TemplateStore
}

func (n *NetHandle) RegType(t interface{}) {
	gob.Register(t)
}

func (n *NetHandle) ExeTmpl(w http.ResponseWriter, templateName string, templateData interface{}) {
	tw := &tmplDataWrapper{templateData, n}
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
		datestring = t.Format("15:04 Jan-02")
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

func NewNet() (nn *NetHandle) {
	nn = new(NetHandle)
	nn.Decoder = schema.NewDecoder()
	nn.Router = NewRouter()
	nn.TemplateStore = NewTemplateStore()

	nn.RegType(new(time.Time))

	return
}

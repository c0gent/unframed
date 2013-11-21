package unframed

import (
	"github.com/gorilla/schema"
	"github.com/nsan1129/auctionLog/log"
	"html/template"
	"net/http"
)

type NetHandle struct {
	Decoder *schema.Decoder
	*Router
	*SessionManager
	templates *template.Template
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

func (n *NetHandle) LoadTemplates(t ...string) {
	n.templates = template.Must(template.ParseFiles(t...))
}

func (n *NetHandle) ExeTmpl(w http.ResponseWriter, templateName string, templateData interface{}) {
	err := n.templates.ExecuteTemplate(w, templateName, templateData)
	if err != nil {
		log.Error(err)
	}
}

func NewNet() (nn *NetHandle) {
	nn = new(NetHandle)

	nn.Decoder = schema.NewDecoder()
	nn.Router = NewRouter()
	return
}

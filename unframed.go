package unframed

import (
	"github.com/nsan1129/auctionLog/log"
	"net/http"
	"strconv"
)

func Atoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Error(err)
	}
	return
}

func QueryUrl(s string, r *http.Request) (i int) {
	i = Atoi(r.URL.Query().Get(s))
	return
}

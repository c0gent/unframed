package unframed

import (
	"github.com/nsan1129/auctionLog/log"
	"strconv"
)

func Atoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Error("unframed.Atoi; string:", s, "int:", i, "][", err)
	}
	return
}

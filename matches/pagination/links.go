package pagination

import (
	"log"
	"strconv"
)

func links(links *Links, addr string, api string, filter string, sort string, order string, start int64, size int64, totalItems int64) Links {
	links.First = build(addr, api, filter, sort, order, start, size)

	return *links
}

// api can either be
func build(addr string, api string, filter string, sort string, order string, start int64, size int64) string {
	if filter != "" {
		filter = "&"
	}
	address := addr + "/" + api + "/?filter=" + filter + "&sort=" + sort + "&order=" + order + "&start=" + strconv.FormatInt(start, 10) + "&size=" + strconv.FormatInt(size, 10)
	log.Println(address)
	return address
}

package main

import (
	"flag"
	"net/http"

	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"
)

var dir = flag.String("dir", "./cache", "path to cached modules files")
var addr = flag.String("listen", "localhost:18888", "service address to listen on")

func main() {
	flag.Parse()
	proxy := goproxy.New()
	proxy.Cacher = &cacher.Disk{Root: *dir}
	http.ListenAndServe(*addr, proxy)
}

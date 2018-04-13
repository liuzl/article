package main

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/crawlerclub/ce"
	"github.com/crawlerclub/dl"
	"github.com/golang/glog"
	"github.com/liuzl/goutil/rest"
	"net/http"
	"strings"
)

var (
	serverAddr = flag.String("addr", ":8080", "bind address")
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("addr=%s  method=%s host=%s uri=%s",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	r.ParseForm()
	debug := false
	debugStr := r.FormValue("debug")
	if debugStr == "on" || debugStr == "true" {
		debug = true
	}
	url := r.FormValue("url")
	req := &dl.HttpRequest{Url: url, Method: "GET", UseProxy: false, Platform: "mobile"}
	res := dl.Download(req)
	if res.Error != nil {
		rest.MustEncode(w, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "error", Message: res.Error.Error()})
		return
	}

	items := strings.Split(res.RemoteAddr, ":")
	ip := ""
	if len(items) > 0 {
		ip = items[0]
	}
	doc := ce.ParsePro(url, res.Text, ip, debug)
	rest.MustEncode(w, doc)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.Handle("/api/", rest.WithLog(ArticleHandler))
	http.Handle("/", http.FileServer(rice.MustFindBox("ui").HTTPBox()))
	glog.Info("server listen on", *serverAddr)
	glog.Error(http.ListenAndServe(*serverAddr, nil))
}

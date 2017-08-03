package main

import (
	"encoding/json"
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/crawlerclub/ce"
	"github.com/crawlerclub/dl"
	"github.com/golang/glog"
	"net/http"
	"strings"
)

var (
	serverAddr = flag.String("addr", ":8080", "bind address")
)

func mustEncode(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		//panic(err)
		e.Encode(err.Error())
	}
}

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
	req := &dl.HttpRequest{Url: url, Method: "GET", UseProxy: false, Platform: "pc"}
	res := dl.Download(req)
	if res.Error != nil {
		mustEncode(w, struct {
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
	mustEncode(w, doc)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.HandleFunc("/api/", ArticleHandler)
	http.Handle("/", http.FileServer(rice.MustFindBox("ui").HTTPBox()))
	glog.Info("server listen on", *serverAddr)
	glog.Error(http.ListenAndServe(*serverAddr, nil))
}

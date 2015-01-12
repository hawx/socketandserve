package main

import (
	"github.com/hawx/serve"

	"log"
	"flag"
	"net"
	"net/http"
	"net/http/httputil"
	"path"
	"strings"
)

var (
	socketDir = flag.String("socket-dir", "", "")
	port      = flag.String("port", "8080", "")
)

func unixDial(sock string) func(string, string) (net.Conn, error) {
	return func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", sock)
	}
}

func director(req *http.Request) {
	req.URL.Scheme = "http"
	req.URL.Host = "dev"
}

func main() {
	flag.Parse()

	if *socketDir == "" {
		log.Fatal("--socket-dir required")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.Host, r.URL)

		if !strings.HasSuffix(r.Host, ".dev") {
			return
		}

		hostName := strings.TrimSuffix(r.Host, ".dev")
		sockPath := path.Join(*socketDir, hostName+".sock")

		proxy := &httputil.ReverseProxy{
			Director:  director,
			Transport: &http.Transport{Dial: unixDial(sockPath)},
		}

		proxy.ServeHTTP(w, r)
	})

	serve.Port(*port, http.DefaultServeMux)
}

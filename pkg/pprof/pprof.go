package pprof

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strconv"
	"time"
)

var listen net.Listener

func Start(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	var err error
	listen, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go http.Serve(listen, mux)
	return nil
}

func Stop() error {
	return listen.Close()
}

// return port
func Auto() int {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	var err error
	rand.Seed(time.Now().Unix())
	port := rand.Intn(65535-4000) + 4000

	for {
		listen, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			break
		}
		port++
	}

	go http.Serve(listen, mux)
	os.WriteFile("pprof.port", []byte(strconv.Itoa(port)), 0666)
	return port
}

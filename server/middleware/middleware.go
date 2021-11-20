package middleware

import (
	"log"
	"net/http"
)

func DoLogging(orgHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		log.Printf("--> [%s:%s] request", req.Method, req.RemoteAddr)
		orgHandler.ServeHTTP(rsp, req)
	})
}

func Check(orgHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		/* check http header */
		ct := req.Header.Get("Content-Type")
		if ct != "application/json" {
			http.Error(rsp, "only support application/json", http.StatusBadRequest)
			return
		}

		orgHandler.ServeHTTP(rsp, req)
	})
}
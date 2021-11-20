package server

import (
	"bookstore/server/middleware"
	"bookstore/store"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type BookServer struct {
	name      string
	instance *http.Server
	db        store.Store
}

func NewBookServer(name string, addr string, db store.Store) *BookServer {
	/* BookServer init */
	srv := &BookServer{
		name: name,
		instance: &http.Server{
			Addr: addr,
		},
		db: db,
	}

	/* path route init */
	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHanler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.getBookHandler).Methods("GET")
	srv.instance.Handler = middleware.DoLogging(middleware.Check(router))

	log.Printf("server info[name:%s, addr:%s]\n", srv.name, srv.instance.Addr)
	return srv
}

func (srv *BookServer) ListenAndServe() (<-chan error, error) {
	errCH := make(chan error)
	go func() {
		log.Printf("====== server start =====\n")
		errCH <- srv.instance.ListenAndServe()
	}()

	select {
	case err := <- errCH:
		return nil, err
	case <-time.After(time.Second):
		return errCH, nil // wait for server starting
	}
}

func (srv *BookServer) Shutdown(ctx context.Context) error {
	return srv.instance.Shutdown(ctx)
}

func (srv *BookServer) createBookHanler(rsp http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(rsp, "decode error: " + err.Error(), http.StatusBadRequest)
		return
	}

	if err := srv.db.CreateBook(&book); err != nil {
		http.Error(rsp, "DB create error: " + err.Error(), http.StatusBadRequest)
		return
	}

	rsp.Write([]byte("OK\n"))
	return
}

func (srv *BookServer) getBookHandler(rsp http.ResponseWriter, req *http.Request) {
	isbn, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(rsp, "Request without ISBN", http.StatusBadRequest)
		return
	}

	book, err := srv.db.GetBook(isbn)
	if err != nil {
		http.Error(rsp, "Get book from DB fail: " + err.Error(), http.StatusBadRequest)
		return
	}

	bookRresponse(rsp, book)
}

func bookRresponse(rsp http.ResponseWriter, book store.Book) {
	body, err := json.Marshal(book)
	if err != nil {
		http.Error(rsp, "Json encode fail: "+ err.Error(), http.StatusInternalServerError)
		return
	}

	rsp.Header().Set("Content-Type", "application/json")
	rsp.Write(body)
}
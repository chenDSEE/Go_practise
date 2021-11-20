package main

import (
	"bookstore/store/factory"
	"fmt"
	"bookstore/server"
	_ "bookstore/internal/store"
)

/**
 * usage:
 * curl -X POST -H "Content-Type:application/json" -d '{"ISBN" : "ISBN-1", "Name" : "book-1", "Authors":["author-1", "author-2"], "Num": 10}' localhost:8080/book
 * curl -X GET -H "Content-Type:application/json" localhost:8080/book/ISBN-1
 */

func main() {
	fmt.Println("====== start =====")
	db, err := factory.New("mem")
	if err != nil {
		panic("main: create store from factory fail")
	}

	srv := server.NewBookServer("server", ":8080", db)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("server stop for ", err)
	}
	fmt.Println("====== stop =====")
}
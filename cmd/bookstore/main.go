package main

import (
	"bookstore/store/factory"
	"context"
	"fmt"
	"bookstore/server"
	_ "bookstore/internal/store"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	errCH, err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("server stop for ", err)
		return
	}

	/* signal handle */
	sigCH := make(chan os.Signal, 1)
	signal.Notify(sigCH, syscall.SIGINT, syscall.SIGTERM)

	/* grace exit */
	select {
	case err = <- errCH:
		log.Println("server fail and exit with:", err)
		return
	case <- sigCH:
		log.Println("server stop for signal")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("server stop and erros:", err.Error())
	}

	fmt.Println("====== stop =====")
}
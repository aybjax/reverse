package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Book struct {
	ID int
	ISBN string
	Author string
	PublishedYear string
}

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}

	defer f.Close()

	log.SetOutput(f)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%q", r.UserAgent())
		log.Printf("%v\n", "aybjax1")
		fmt.Printf("%v\n", "aybjax2")

		book := Book{
			ID: 123,
			ISBN: "0-201-03801-3",
			Author: "Donald Knuth",
			PublishedYear: "1968",
		}

		jsonData, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	s := &http.Server{
		Addr: ":8080",
		ReadTimeout: time.Second * 15,
		WriteTimeout: time.Second * 15,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
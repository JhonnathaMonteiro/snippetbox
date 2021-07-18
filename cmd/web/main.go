package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "\033[36mINFO\033[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "\033[31mERROR\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	filerServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", filerServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatalln(err)
}

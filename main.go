package main

import (
	"flag"
	"io"
	"net/http"
)

var (
	port string
)

func main() {

	flag.StringVar(&port, "port", "30000", "待ち受けポート番号")
	flag.Parse()

	http.HandleFunc("/", endpoint)
	http.ListenAndServe(":"+port, nil)
}

func endpoint(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")

	switch r.Method {
	case "GET":
		val := get(key)
		if val == "" {
			http.Error(w, "Key Not Found : "+key, http.StatusNotFound)
		}
		io.WriteString(w, val)
	case "POST", "PUT":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Sever Error : Request ParseError", http.StatusBadRequest)
		}
		key := r.FormValue("key")
		value := r.FormValue("value")
		if key == "" || value == "" {
			http.Error(w, "Sever Error : key "+key+" value "+value, http.StatusBadRequest)
		}
		postOrPut(key, value)
	case "DELETE":
		del(key)
	default:
		http.Error(w, "Doesn't support protocol.", http.StatusNotAcceptable)
	}
}

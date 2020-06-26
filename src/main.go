package main

import "log"
import "net/http"
import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/cars", GetCars)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}

func GetCars(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`["BMW","Mercedes","Bentley","Audi","Cadillac"]`))
}

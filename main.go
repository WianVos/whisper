package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/check", CheckHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("done")
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	//now run the healthcheck
	w.Header().Set("Content-Type", "application/json")

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}

}

// func getJSONResponse(s string) []byte {
// 	result := struct {
// 		Status string
// 	}{
// 		Status: s,
// 	}

// 	payload, _ := json.Marshal(result)

// 	return payload
// }

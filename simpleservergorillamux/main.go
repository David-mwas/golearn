// go install -v golang.org/x/tools/gopls@latest
package simpleservergorillamux

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

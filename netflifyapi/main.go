package main

import (
	"fmt"
	"net/http"

	"github.com/David-mwas/golearn/routers"
)

func main() {

	fmt.Println("hello gophers...")
	r := routers.Router()

	http.ListenAndServe(":5000", r)
	fmt.Println("sever listening at port 5000...")

}

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	//fmt.Println("main test 입니다.")

	router := httprouter.New()
	router.GET("/hello", HelloAPI)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func HelloAPI(rs http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(rs, "Hello GO!")
}

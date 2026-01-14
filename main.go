/*PACKAGE*/
package main

/**/

/*IMPORTS FROM PACKAGE*/
import (
	"log"
	"net/http"
	WebController "seachProfile/controller"

	"github.com/gorilla/mux"
)

/**/

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/web", WebController.GetMostPopularWebs).Methods("GET")
	router.HandleFunc("/web", WebController.CreateUsersWebs).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}

/*PACKAGE*/
package WebController

/**/

/*IMPORTS FROM PACKAGE*/
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"seachProfile/models"
	webService "seachProfile/service"
)

/**/

func CreateUsersWebs(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	fmt.Println("IMPRIMIENTO FROM CONTROLLER")
	var web models.Webs
	json.NewDecoder(r.Body).Decode(&web)
	fmt.Println("JSON recibido:")
	fmt.Printf("%+v\n", web)

	status, err := webService.CreateUsersService(ctx, web)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(status)
}

func GetMostPopularWebs(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	webs, err := webService.GetUsersWebsService(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(webs)
}

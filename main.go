/*PACKAGE*/
package main

/**/

/*IMPORTS FROM PACKAGE*/
import (
	// "log" local

	//"github.com/gorilla/mux" local
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/**/

/*LOCAL*/
// func main() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/web", WebController.GetMostPopularWebs).Methods("GET")
// 	router.HandleFunc("/web", WebController.CreateUsersWebs).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":3000", router))
// }

/*END*/

/*LAMBDA HANDLER*/

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Lambda Go funcionando 🚀",
	}, nil
}

func main() {
	lambda.Start(Handler)
}

/*END*/

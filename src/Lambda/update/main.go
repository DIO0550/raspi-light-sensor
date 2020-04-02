// main.go
package main

import ( 
	"log"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/Lambda/package/db"
)

const (
	HttpError500NoRequiredParameters string = "500 No required parameters"
)

func update(event db.UpdateConferenceRoomParam) (db.UpdateConferenceRoomResult, error) {
	name := event.Name
	isUse := event.IsUse
	fmt.Println(event)
	if (name == nil || isUse == nil) {
		log.Printf(HttpError500NoRequiredParameters)
		result := db.UpdateConferenceRoomResult{
			ResultCode: "500",
			Message: HttpError500NoRequiredParameters,
		}
		return result, nil
	}
	result := db.UpdateConferenceRoomData(*name, *isUse)
	fmt.Println(result)
	return result, nil
}

func main() {
	lambda.Start(update)
}


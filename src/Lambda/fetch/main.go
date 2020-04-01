// main.go
package main

import ( 
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/Lambda/package/db"
)

func fetch() ([]db.ConferenceRoom, error) {
	rooms, err := db.FetchAllConferenceRoomData()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	return rooms, nil
}

func main() {
	lambda.Start(fetch)
}

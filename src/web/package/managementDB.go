package  managementDB

import (
	"database/sql"
    "fmt"
    "log"
    
    _ "github.com/go-sql-driver/mysql"
)

type ConferenceRoom struct {
    Id int  `json:"id"`
    Name string `json:"name"`
    UsageSituation bool `json:"usage_situation"`
}

type UpdateConferenceRoomParam struct {
    Name *string `json:"name"`
    UsageSituation *bool `json:"usage_situation"`
}

func ConnectionDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "golang:golang@tcp(192.168.0.2:3306)/golang?parseTime=true")
    if err != nil {
        log.Printf(err.Error())
        fmt.Println(err)
    }
    return db, err
}

func FetchAllConferenceRoomData()(conferenceRooms []ConferenceRoom, err error) {
    db, err := ConnectionDB()
    if err != nil {
        log.Printf(err.Error())
        fmt.Println(err)
        return nil, err;
    } 
    defer db.Close()

    rows, err := db.Query("SELECT * FROM conference_room ORDER BY NAME ASC")
    if err != nil {
        log.Printf(err.Error())
        fmt.Println(err)
        return nil, err;
    }

    for rows.Next() {
        r := ConferenceRoom{}
        err = rows.Scan(&r.Id, &r.Name, &r.UsageSituation)
        if err != nil {
            log.Printf(err.Error())
            fmt.Println(err)
            return nil, err;
        }
        conferenceRooms = append(conferenceRooms, r)
    }
    
    return
}

func UpdateConferenceRoomData(roomName string) {

}
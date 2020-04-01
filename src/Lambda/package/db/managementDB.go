package db

import (
	"database/sql"
    "fmt"
    "log"
    "strconv"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

type ConferenceRoom struct {
    Id int  `json:"id"`
    Name string `json:"name"`
    IsUse bool `json:"is_use"`
}

/// ConferenceRoom更新パラメータ
type UpdateConferenceRoomParam struct {
    Name *string `json:"name"`
    IsUse *bool `json:"is_use"`
}

/// ConferenceRoom更新結果
type UpdateConferenceRoomResult struct {
    ResultCode string `json:"result_code"`
    Message string `json:"message"`
}

func ConnectionDB() (*sql.DB, error) {
    DBName :=  os.Getenv("DataBaseName")
    User := os.Getenv("User")
    Password := os.Getenv("Password")
    Port := os.Getenv("Port")
    EntryPoint := os.Getenv("EntryPoint")
    connection := User + ":" + Password + "@tcp(" + EntryPoint + ":" + Port + ")" + "/" + DBName
    // user:password@tcp(entrypoint:port)/dbname?
    db, err := sql.Open("mysql", connection)
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
        err = rows.Scan(&r.Id, &r.Name, &r.IsUse)
        if err != nil {
            log.Printf(err.Error())
            fmt.Println(err)
            return nil, err
        }
        conferenceRooms = append(conferenceRooms, r)
    }

    return
}

func UpdateConferenceRoomData(roomName string, IsUse bool) (updateResult UpdateConferenceRoomResult)  {
    db, err := ConnectionDB()
    if err != nil {
        log.Printf(err.Error())
        fmt.Println(err)
        updateResult.ResultCode = "001"
        updateResult.Message = err.Error()
        return
    }
    defer db.Close()
    result , err := db.Exec("UPDATE conference_room SET usage_situation = ? WHERE name = ?", IsUse, roomName)
    if err != nil {
        updateResult.ResultCode = "002"
        updateResult.Message = err.Error()
        return
    }

    id, _ := result.RowsAffected()
    updateResult.ResultCode = "000"
    updateResult.Message = "chage id=" + string(id) + " value=" + strconv.FormatBool(IsUse)
    return
}

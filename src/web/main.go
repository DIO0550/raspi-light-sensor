package main

import (
	"html/template"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/web/package/db"
	"github.com/web/package/converter"
)

var templates = make(map[string]*template.Template)

const  (
	HttpError405 string = "405 Method Not Allowed"
	HttpError415 string = "415 Unsupported Media Type"
	HttpError500ParameterFormatError string = "500 Parameter format error"
	HttpError500NoRequiredParameters string = "500 No required parameters"

	AllowOrigin = "*"
)

type TemplateData struct {
	Title string
	Data interface{}
	DataMap map[string][]db.ConferenceRoom
}


func main() {
	port := "8080"
	templates["index"] = LoadTemplate("index")

	// css読み込み
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", http.FileServer(http.Dir("stylesheet/"))))

	// root
	http.HandleFunc("/", HandleIndex)

	// conferenceroom
	http.HandleFunc("/api/conferenceroom/update", HandleConferenceRoomUpdate)
	http.HandleFunc("/api/conferenceroom/fetch", HandleConferenceRoomFetch)

	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}


// テンプレートをロードする
func LoadTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles("template/"+templateName+".html", "template/_header.html", "template/_footer.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	return t
}


// indexのハンドルを返す
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// 存在しないパスなら、404を返す
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	rooms, err := db.FetchAllConferenceRoomData()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}


	data := TemplateData{Title: "index", Data: rooms, DataMap: map[string][]db.ConferenceRoom{}}

	for _, value := range rooms {
		name := value.Name
		number := converter.ExtractNumber(name)
		log.Print(name)
		log.Print(number)
		dataMap := data.DataMap[string(number)]
		if dataMap == nil {
			log.Print("data_map == nil")
			emptySlice := []db.ConferenceRoom{}
			emptySlice = append(emptySlice, value)
			data.DataMap[string(number)] = emptySlice
		} else {
			log.Print("data_map != nil")
			dataMap = append(dataMap, value)
			data.DataMap[string(number)] = dataMap
		}
	}

	if err = templates["index"].Execute(w, data)
	 err != nil {
		http.Error(w, err.Error(), 500)
	}
}


// ConfrenceRoomUpdateのハンドルを返す
func HandleConferenceRoomUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, HttpError405, http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, HttpError415, http.StatusUnsupportedMediaType)
		return
	}


	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 不要なキーの確認
	for key, _ := range m {
		if key != "name" && key != "usage_situation" {
			http.Error(w, HttpError500ParameterFormatError, 500)
			return
		}
	}


	var room db.UpdateConferenceRoomParam
	err = json.Unmarshal(b, &room)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 必須パラメータチェック
	if room.Name == nil || room.UsageSituation == nil {
		http.Error(w, HttpError500NoRequiredParameters, 500)
		return
	}

	updateResult := db.UpdateConferenceRoomData(*room.Name, *room.UsageSituation)
	jsonRespose, _ := json.Marshal(updateResult)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRespose)
}


// ConfrenceRoomFetchのハンドルを返す
// DBから取得した会議室の在室状況を返却する
func HandleConferenceRoomFetch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, HttpError405, http.StatusMethodNotAllowed)
		return
	}

	rooms, err := db.FetchAllConferenceRoomData()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	jsonRespose, _ := json.Marshal(rooms)
	SetUpCORS(w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRespose)
}

// CORS設定
func SetUpCORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", AllowOrigin)
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
}
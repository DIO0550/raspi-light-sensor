package main

import (
	"html/template"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/web/package"
)

var templates = make(map[string]*template.Template)

const  (
	HttpError405 string = "405 Method Not Allowed"
	HttpError415 string = "415 Unsupported Media Type"
	HttpError500ParameterFormatError string = "500 Parameter format error"
	HttpError500NoRequiredParameters string = "500 No required parameters"
)

type TemplateData struct {
	Title string
	Data interface{}
}

func main() {
	port := "8080"
	templates["index"] = LoadTemplate("index")

	// css読み込み
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", http.FileServer(http.Dir("stylesheet/")))) 
	
	// root
	http.HandleFunc("/", HandleIndex)
	
	// conferenceroom
	http.HandleFunc("api/conferenceroom/update", HandleConferenceRoomUpdate)

	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))

}

// indexのハンドルを返す
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// 存在しないパスなら、404を返す
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return 
	}
	rooms, err := managementDB.FetchAllConferenceRoomData()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	data := TemplateData{Title: "index", Data: rooms}

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


	var room managementDB.UpdateConferenceRoomParam
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

	updateResult := managementDB.UpdateConferenceRoomData(*room.Name, *room.UsageSituation)
	jsonRespose, _ := json.Marshal(updateResult)
	w.Header().Set("Content-Type", "application/json")
    w.Write(jsonRespose)
}

// テンプレートをロードする
func LoadTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles("template/"+templateName+".html", "template/_header.html", "template/_footer.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	return t
}
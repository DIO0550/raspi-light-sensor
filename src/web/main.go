package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/web/package"
)

var templates = make(map[string]*template.Template)

type TemplateData struct {
	Title string
	Data interface{}
}

func main() {
	port := "8080"
	templates["index"] = LoadTemplate("index")
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", http.FileServer(http.Dir("stylesheet/")))) 
	http.HandleFunc("/", HandleIndex)
	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

// indexのハンドルを返す
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	rooms, err := managementDB.FetchAllConferenceRoomData()
	if err != nil {
		log.Printf(err.Error())
	}

	data := TemplateData{Title: "index", Data: rooms}

	if err = templates["index"].Execute(w, data)
	 err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

// テンプレートをロードする
func LoadTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles("template/"+templateName+".html", "template/_header.html", "template/_footer.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	return t
}
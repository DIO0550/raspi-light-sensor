package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = make(map[string]*template.Template)

func main() {
	port := "8080"
	templates["index"] = loadTemplate("index")
	http.HandleFunc("/", handleIndex)
	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := templates["index"].Execute(w, nil)
	 err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func loadTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles("template/"+templateName+".html", "template/_header.html", "template/_footer.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	return t
}
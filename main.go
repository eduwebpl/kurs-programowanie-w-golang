package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	assetDirectory := http.FileServer(http.Dir("./static/assets"))
	imagesDirectory := http.FileServer(http.Dir("./static/images"))

	http.Handle("/assets/", http.StripPrefix("/assets", assetDirectory))
	http.Handle("/images/", http.StripPrefix("/images", imagesDirectory))
	http.HandleFunc("/", rootPath)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type SiteTemplate struct {
	Title string
	Name  string
}

func rootPath(writer http.ResponseWriter, req *http.Request) {
	siteData := SiteTemplate{
		"Eduweb",
		"Kamil",
	}

	tmplt, err := template.ParseFiles("./static/index.html")
	if err != nil {
		return
	}

	tmplt.Execute(writer, siteData)
}

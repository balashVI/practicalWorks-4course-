package main

import (
	"html/template"
	"net/http"
)

func patternRLab1(writer http.ResponseWriter, request *http.Request) {
	patternRLab1Page := template.Must(template.ParseFiles("tmpl/pageLayout.html", "tmpl/patternRLab1.html"))
	patternRLab1Page.Execute(writer, Page{"Розпізнавання образів Лаб. 1",
		"Розпізнавання образів Лаб. 1", []string{"/js/patternRLab1.js"}, nil})
}

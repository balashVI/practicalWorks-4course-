package main

import (
	"html/template"
	"net/http"
)

func main() {
	//Надання достап до файлів на сервері
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//Реєстрація обробників
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/DSP/lab1", DSPLab1)
	http.HandleFunc("/DSP/lab1Calc", DSPLab1Calc)
	http.HandleFunc("/DSP/lab2", DSPLab2)
	http.HandleFunc("/DSP/lab2Calc", DSPLab2Calc)

	//Запуск сервера
	http.ListenAndServe(":8081", nil)
}

func mainPage(writer http.ResponseWriter, request *http.Request) {
	menuPage := template.Must(template.ParseFiles("tmpl/pageLayout.html", "tmpl/mainMenu.html"))
	menuPage.Execute(writer, Page{"Балаш Віталій ФеІ-41",
		"Лабораторні роботи студента групи ФеІ-41 Балаша Віталія", nil, lessons})
}

type Page struct {
	Title      string
	PageHeader string
	JSLinks    []string
	Content    interface{}
}

var lessons = []lesson{
	{
		"Цифрова обробка сигналів",
		[]lab{
			{"Лабораторна робота №1", "/DSP/lab1"},
			{"Лабораторна робота №2", "/DSP/lab2"}}},
	{
		"Розпізнаванн образів",
		[]lab{
			{"Лабораторна робота №1", "#"},
			{"Лабораторна робота №2", "#"}}}}

type lab struct {
	Name string
	Link string
}

type lesson struct {
	Name string
	Labs []lab
}

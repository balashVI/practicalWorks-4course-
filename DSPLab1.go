package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func DSPLab1(writer http.ResponseWriter, request *http.Request) {
	DSPLab1Page := template.Must(template.ParseFiles("tmpl/pageLayout.html", "tmpl/DSPLab1.html"))
	DSPLab1Page.Execute(writer, Page{"Цифрова обробка сигналів Лаб. 1",
		"Цифрова обробка сигналів Лаб. 1", nil, nil})
}

func DSPLab1Calc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	vec1str := strings.Split(request.URL.Query()["vec1"][0], " ")
	vec2str := strings.Split(request.URL.Query()["vec2"][0], " ")
	var vec1, vec2 []int

	var number int
	var err error
	for _, i := range vec1str {
		number, err = strconv.Atoi(i)
		if err != nil {
			fmt.Fprintln(writer, "Введені некоректні дані")
			return
		}
		vec1 = append(vec1, number)
	}
	for _, i := range vec2str {
		number, err = strconv.Atoi(i)
		if err != nil {
			fmt.Fprintln(writer, "Введені некоректні дані")
			return
		}
		vec2 = append(vec2, number)
	}

	fmt.Fprintf(writer, "Перший вектор: %v\nДругий вектор: %v\n", vec1, vec2)

	res := make([]int, len(vec1)+len(vec2)-1)
	for m := range res {
		for k := range vec1 {
			if m-k >= 0 && m-k < len(vec2) {
				res[m] += vec1[k] * vec2[m-k]
			}
		}
	}
	fmt.Fprintf(writer, "Результат: %v\n", res)
}

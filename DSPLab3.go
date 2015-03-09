package main

import (
	"encoding/json"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func DSPLab3(writer http.ResponseWriter, request *http.Request) {
	DSPLab3Page := template.Must(template.ParseFiles("tmpl/pageLayout.html", "tmpl/DSPLab3.html"))
	DSPLab3Page.Execute(writer, Page{"Цифрова обробка сигналів Лаб. 3 (ДПФ)",
		"Цифрова обробка сигналів Лаб. 3 (ДПФ)", []string{"/js/angular.min.js", "/js/DSPLab3.js"}, nil})
}

func DSPLab3Calc(writer http.ResponseWriter, request *http.Request) {
	mode, _ := strconv.Atoi(request.FormValue("mode"))
	var inputData []Complex
	err := json.Unmarshal([]byte(request.FormValue("inputData")), &inputData)
	if err != nil {
		http.Error(writer, err.Error(), 0)
	}
	resData, _ := json.Marshal(dft(inputData, mode))
	writer.Write(resData)
}

func dft(data []Complex, inv int) []Complex {
	n := len(data)
	res := make([]Complex, n)
	WN := 2.0 * math.Pi / float64(n)
	if inv == 1 {
		WN = -WN
	}
	var wk, c, s float64
	for i := 0; i < n; i++ {
		wk = float64(i) * WN
		for j := range data {
			c = math.Cos(float64(j) * wk)
			s = math.Sin(float64(j) * wk)
			res[i].Re += data[j].Re*c + data[j].Im*s
			res[i].Im -= data[j].Re*s + data[j].Im*c
		}
		if inv == 1 {
			res[i].Re /= float64(n)
			res[i].Im /= float64(n)
		}
	}
	return res
}

type Complex struct {
	Re, Im float64
}

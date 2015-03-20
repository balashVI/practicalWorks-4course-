package main

import (
	"encoding/json"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"fmt"
)

func DSPLab3(writer http.ResponseWriter, request *http.Request) {
	DSPLab3Page := template.Must(template.ParseFiles("tmpl/pageLayout.html", "tmpl/DSPLab3.html"))
	DSPLab3Page.Execute(writer, Page{"Цифрова обробка сигналів Лаб. 3 (ДПФ)",
		"Цифрова обробка сигналів Лаб. 3 (ДПФ)", []string{"/js/angular.min.js", "/js/DSPLab3.js", "/js/Chart.js"}, nil})
}

func DSPLab3Calc2(writer http.ResponseWriter, request *http.Request) {
	T := 1.0
	N := 30
	inpSignal1 := make([]Complex, N)
	inpSignal2 := make([]Complex, N)
	time := make([]string, N)
	frequency := make([]float64, N)

	dt := T / float64(N)
	for i := 0; i < N; i++ {
		inpSignal1[i].Re = func1(dt*float64(i), T)
		inpSignal2[i].Re = func2(dt*float64(i), T)
		time[i] = fmt.Sprintf("%.6f", dt * float64(i))
		frequency[i] =  float64(i)/(dt*float64(N))
	}
	resSignal1 := dft(inpSignal1, 0)
	resSignal2 := dft(inpSignal2, 0)

	resData := struct {
		Time []string
		Frequency []float64
		InpSignal1, ResSignal1, InpSignal2, ResSignal2 []Complex
	}{time, frequency,inpSignal1, resSignal1, inpSignal2, resSignal2}
	resJSON, _ := json.Marshal(resData)
	writer.Write(resJSON)
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

func func1(x, T float64) float64 {
	x = math.Mod(x, T)
	if x < T/2.0 {
		return 1
	}
	return 0
}

func func2(x, T float64) float64 {
	x = math.Mod(x, T)
	var a, b float64
	a = 2.0 / T
	if x >= T/2.0 {
		a = -a
		b = 2
	}
	return a*x + b
}

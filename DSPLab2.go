package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func DSPLab2(writer http.ResponseWriter, request *http.Request) {
	DSPLab1Page := template.Must(template.ParseFiles("tmpl/pageLayout.html", "tmpl/DSPLab2.html"))
	DSPLab1Page.Execute(writer, Page{"Цифрова обробка сигналів Лаб. 2",
		"Цифрова обробка сигналів Лаб. 2", []string{"/js/DSPLab2.js", "/js/Chart.js"}, nil})
}

func DSPLab2Calc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	vec1str := strings.Split(request.URL.Query()["vec1"][0], " ")
	vec2str := strings.Split(request.URL.Query()["vec2"][0], " ")
	scaling, err := strconv.Atoi(request.URL.Query()["scaling"][0])
	if err != nil {
		fmt.Fprintln(writer, "Введені некоректні дані")
		return
	}
	shift, err := strconv.Atoi(request.URL.Query()["shift"][0])
	if err != nil {
		fmt.Fprintln(writer, "Введені некоректні дані")
		return
	}
	timeScaling, err := strconv.Atoi(request.URL.Query()["timeScaling"][0])
	if err != nil {
		fmt.Fprintln(writer, "Введені некоректні дані")
		return
	}
	var vec1, vec2 []int
	var number int
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

	resData := struct {
		Vector1        []int
		Vector2        []int
		Scaling        int
		Shift          int
		TimeScaling    int
		ScalingRes     []int
		ReversRes      []int
		ShiftRes       []int
		TimeScalingRes []int
		AddRes         []int
		MulRes         []int
	}{vec1, vec2, scaling, shift, timeScaling, DSPLab2Scaling(vec1, scaling),
		DSPLab2Revers(vec1), DSPLab2Shift(vec1, shift), DSPLab2TimeScaling(vec1, timeScaling),
		DSPLab2Add(vec1, vec2), DSPLab2Mul(vec1, vec2)}

	resJson, err := json.Marshal(resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	writer.Write(resJson)

	//	outBufer := bufio.NewWriter(writer)

	//	fmt.Fprintln(outBufer, "Перший сигнал: ", vec1)
	//	fmt.Fprintln(outBufer, "Другий сигнал: ", vec2)
	//	fmt.Fprintln(outBufer, "Коефіцієнт масштабування: ", scaling)
	//	fmt.Fprintln(outBufer, "Коефіцієнт зсуву: ", shift)
	//	fmt.Fprintln(outBufer, "Коефіцієнт розширення: ", timeScaling)
	//	fmt.Fprintln(outBufer, "Масштабування першого сигнала: ", DSPLab2Scaling(vec1, scaling))
	//	fmt.Fprintln(outBufer, "Реверс першого сигнала: ", DSPLab2Revers(vec1))
	//	fmt.Fprintln(outBufer, "Зсув по часу першого сигнала: ", DSPLab2Shift(vec1, shift))
	//	fmt.Fprintln(outBufer, "Зсув по часу першого сигнала: ", DSPLab2TimeScaling(vec1, timeScaling))
	//	fmt.Fprintln(outBufer, "Додавання сигналів: ", DSPLab2Add(vec1, vec2))
	//	fmt.Fprintln(outBufer, "Множення сигналів: ", DSPLab2Mul(vec1, vec2))

	// outBufer.Flush()
}

func DSPLab2Revers(inp []int) []int {
	n := len(inp)
	out := make([]int, n)
	for i := range inp {
		out[i] = inp[n-i-1]
	}
	return out
}

func DSPLab2Scaling(inp []int, scaling int) []int {
	out := make([]int, len(inp))
	for i := range inp {
		out[i] = scaling * inp[i]
	}
	return out
}

func DSPLab2TimeScaling(inp []int, timeScaling int) []int {
	out := make([]int, len(inp)*timeScaling)
	for i := range inp {
		out[i*timeScaling] = inp[i]
		if i+1 < len(inp) {
			for j := 1; j < timeScaling; j += 1 {
				out[i*timeScaling+j] = (inp[i] + inp[i+1]) * (timeScaling - j) / timeScaling
			}
		}
	}
	for j := 1; j < timeScaling; j += 1 {
		out[(len(inp)-1)*timeScaling+j] = inp[len(inp)-1] * (timeScaling - j) / timeScaling
	}
	return out
}

func DSPLab2Shift(inp []int, shift int) []int {
	out := make([]int, len(inp)+shift)
	for i := range inp {
		out[i+shift] = inp[i]
	}
	return out
}

func DSPLab2Add(inp1, inp2 []int) []int {
	n1 := len(inp1)
	n2 := len(inp2)
	var n int
	if n1 > n2 {
		n = n1
	} else {
		n = n2
	}
	res := make([]int, n)
	for i := range inp1 {
		res[i] = inp1[i]
	}
	for i := range inp2 {
		res[i] += inp2[i]
	}
	return res
}

func DSPLab2Mul(inp1, inp2 []int) []int {
	n1 := len(inp1)
	n2 := len(inp2)
	var n int
	if n1 > n2 {
		n = n1
	} else {
		n = n2
	}
	res := make([]int, n)
	for i := range res {
		if i < n1 && i < n2 {
			res[i] = inp1[i] * inp2[i]
		} else {
			break
		}
	}
	return res
}

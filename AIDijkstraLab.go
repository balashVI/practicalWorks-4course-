package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/http"
	textTemplate "text/template"
)

func AIDijkstraLab(writer http.ResponseWriter, request *http.Request) {
	AIDijkstraLabPage := template.Must(template.ParseFiles("tmpl/pageLayout.html", "tmpl/AIDijkstraLab.html"))
	AIDijkstraLabPage.Execute(writer, Page{"Штучний інтелект Лаб. 1 (алгоритм Дейкстри)",
		"Штучний інтелект Лаб. 1 (алгоритм Дейкстри)", []string{"/js/angular.min.js", "/js/AIDijkstraLab.js"}, nil})
}

var inf = -1

func AIDijkstraLabCalc(writer http.ResponseWriter, request *http.Request) {
	//Валідація даних
	var requestData struct {
		StartPoint, FinishPoint, NumberOfVertices int
		EdgesMatrix                               []int
	}
	json.Unmarshal([]byte(request.FormValue("requestData")), &requestData)
	edgesMatrix := make([][]int, requestData.NumberOfVertices)
	for i := range edgesMatrix {
		edgesMatrix[i] = make([]int, requestData.NumberOfVertices)
		for j := range edgesMatrix {
			if i == j {
				edgesMatrix[i][j] = -1
			} else if i < j {
				edgesMatrix[i][j] = requestData.EdgesMatrix[i*requestData.NumberOfVertices+j]
			} else {
				edgesMatrix[i][j] = requestData.EdgesMatrix[j*requestData.NumberOfVertices+i]
			}
		}
	}
	if requestData.StartPoint < 0 || requestData.StartPoint >= requestData.NumberOfVertices {
		requestData.StartPoint = 0
	}
	if requestData.FinishPoint < 0 || requestData.FinishPoint >= requestData.NumberOfVertices {
		requestData.FinishPoint = requestData.NumberOfVertices - 1
	}

	distance, path := dijkstra(requestData.StartPoint, requestData.FinishPoint, edgesMatrix)
	var gr Graph
	gr.CreateCircles(requestData.NumberOfVertices)
	gr.CreateLines(edgesMatrix)
	slides := gr.GenerateSlides(path)

	resData := struct {
		StartPoint, FinishPoint int
		EdgesMatrix             []int
		Distance                int
		Path                    []int
		Slides                  []string
	}{
		StartPoint:  requestData.StartPoint,
		FinishPoint: requestData.FinishPoint,
		Distance:    distance,
		Path:        path,
		Slides:      slides}
	for i := range edgesMatrix {
		resData.EdgesMatrix = append(resData.EdgesMatrix, edgesMatrix[i]...)
	}

	bytesRes, err := json.Marshal(resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	writer.Write(bytesRes)
}

func dijkstra(startPoint, finishPoint int, graph [][]int) (int, []int) {
	path := make([]int, 0)
	path = append(path, startPoint)
	distances := make([]int, len(graph))
	for i := range distances {
		distances[i] = inf
	}
	distances[startPoint] = 0
	visited := make([]bool, len(graph))

	currentPoint := startPoint
	for currentPoint != finishPoint {
		visited[currentPoint] = true
		//Оновлюємо відстані
		for i := range visited {
			if visited[i] || graph[currentPoint][i] == inf {
				continue
			}

			var newDistance int
			if distances[currentPoint] == inf {
				distances[i] = graph[currentPoint][i]
			} else {
				newDistance = distances[currentPoint] + graph[currentPoint][i]
				if newDistance < distances[i] {
					distances[i] = newDistance
				}
			}
		}

		//Шукаємо наступну точку
		nextPoint := -1
		for i := range visited {
			if visited[i] || graph[currentPoint][i] == inf {
				continue
			}
			if nextPoint == -1 {
				nextPoint = i
			}
			if distances[nextPoint] > distances[i] {
				nextPoint = i
			}
		}
		if nextPoint == -1 {
			return -1, path
		}
		path = append(path, nextPoint)
		currentPoint = nextPoint
	}
	return distances[finishPoint], path
}

type point struct {
	X, Y float64
}

type line struct {
	StartPoint, EndPoint point
	StrokeWidth          int
	StrokeColor          string
}

type circle struct {
	Number      int
	CenterPoint point
	Radius      float64
	StrokeWidth int
	StrokeColor string
	FillColor   string
}

type Graph struct {
	Lines   map[int]*line
	Circles []circle
}

func (gr *Graph) CreateCircles(circlesN int) {
	var angleStep float64
	canvasCenter := 500.0
	angleStep = 2.0 * math.Pi / float64(circlesN)
	r := 350.0
	сirclesRadius := math.Pi * r / float64(circlesN) * 0.5

	gr.Circles = make([]circle, circlesN)
	var currentAngle float64
	for i := 0; i < circlesN; i++ {
		gr.Circles[i].Number = i
		gr.Circles[i].CenterPoint.X = canvasCenter + r*math.Cos(currentAngle)
		gr.Circles[i].CenterPoint.Y = canvasCenter + r*math.Sin(currentAngle)
		gr.Circles[i].Radius = сirclesRadius
		gr.Circles[i].StrokeWidth = 4
		gr.Circles[i].StrokeColor = "#4B4B4B"
		gr.Circles[i].FillColor = "#FAFAFA"
		currentAngle += angleStep
	}
}

func (gr *Graph) CreateLines(matrix [][]int) {
	gr.Lines = make(map[int]*line)
	n := len(matrix)
	for i := range matrix {
		for j := i; j < n; j++ {
			if matrix[i][j] != inf {
				gr.Lines[n*i+j] = &line{gr.Circles[i].CenterPoint, gr.Circles[j].CenterPoint, 4, "#4B4B4B"}
			}
		}
	}
}

func (gr *Graph) GenerateSlides(path []int) []string {
	fmt.Println(path)
	res := make([]string, len(path))

	strokeColor := "#E60000"
	fillColor := "#FEFAFA"

	gr.Circles[path[0]].StrokeColor = "#009933"
	gr.Circles[path[0]].FillColor = "#FAFCFA"
	gr.Circles[path[0]].StrokeWidth = 8
	gr.Circles[len(path)-1].StrokeColor = "#009933"
	gr.Circles[len(path)-1].FillColor = "#FAFCFA"
	gr.Circles[len(path)-1].StrokeWidth = 8

	tmpl := textTemplate.Must(textTemplate.ParseFiles("tmpl/AIDijkstraLab.svg"))

	var wr bytes.Buffer

	tmpl.Execute(&wr, gr)
	res[0] = wr.String()

	for i := 1; i < len(path)-1; i++ {
		wr.Reset()
		var lineN int
		if path[i-1] < path[i] {
			lineN = path[i-1]*len(gr.Circles) + path[i]
		} else {
			lineN = path[i]*len(gr.Circles) + path[i-1]
		}
		gr.Lines[lineN].StrokeColor = strokeColor
		gr.Lines[lineN].StrokeWidth = 8
		gr.Circles[path[i]].StrokeColor = strokeColor
		gr.Circles[path[i]].FillColor = fillColor
		gr.Circles[path[i]].StrokeWidth = 8
		tmpl.Execute(&wr, gr)
		res[i] = wr.String()
	}

	wr.Reset()
	var lineN int
	lastIndex := len(path) - 1
	if path[lastIndex-1] < path[lastIndex] {
		lineN = path[lastIndex-1]*len(gr.Circles) + path[lastIndex]
	} else {
		lineN = path[lastIndex]*len(gr.Circles) + path[lastIndex-1]
	}
	gr.Lines[lineN].StrokeColor = strokeColor
	gr.Lines[lineN].StrokeWidth = 6
	tmpl.Execute(&wr, gr)
	res[len(path)-1] = wr.String()

	return res
}

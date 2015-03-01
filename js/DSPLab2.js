document.onLoad = function () {
    console.log("hello")
}

function sendData() {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function () {
        if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
            var res = JSON.parse(xmlhttp.responseText)
            document.getElementById("res").innerHTML = "Перший сигнал: " + res.Vector1 +
                "\nДругий сигнал: " + res.Vector2 + "\nКоефіцієнт масштабування: " + res.Scaling +
                "\nКоефіцієнт зсуву: " + res.Shift + "\nКоефіцієнт розширення: " + res.TimeScaling +
                "\nМасштабування першого сигнала: " + res.ScalingRes +
                "\nРеверс першого сигнала: " + res.ReversRes +
                "\nЗсув першого сигнала: " + res.ShiftRes +
                "\nРозширення першого сигнала: " + res.TimeScalingRes +
                "\nДодавання сигналів: " + res.AddRes +
                "\nМноження сигналів: " + res.MulRes;

            createChart(res.Vector1, "vec1chart")
            createChart(res.Vector2, "vec2chart")
            createChart(res.ScalingRes, "scalingChart")
            createChart(res.ReversRes, "reversChart")
            createChart(res.ShiftRes, "shiftChart")
            createChart(res.TimeScalingRes, "timeScalingChart")
            createChart(res.AddRes, "addChart")
            createChart(res.MulRes, "mulChart")
        }
    }
    xmlhttp.open("GET", "/DSP/lab2Calc?vec1=" + document.getElementById("vec1").value +
        ";vec2=" + document.getElementById("vec2").value + ";scaling=" +
        document.getElementById("scaling").value + ";shift=" +
        document.getElementById("shift").value + ";timeScaling=" +
        document.getElementById("timeScaling").value, true);
    xmlhttp.send();
}

function createChart(dataVector, canvasId) {
    var labelsList = Array(dataVector.length)
    for (var i = 0; i < dataVector.length; i++)
        labelsList[i] = i.toString()

    new Chart(document.getElementById(canvasId).getContext("2d")).Line({
        labels: labelsList,
        datasets: [
            {
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: dataVector
        }
    ]
    }, null);

}
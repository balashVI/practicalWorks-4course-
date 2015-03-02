var categories = [
    [4, 2, 0, 0],
    [2, 0, 1, 0],
    [1, 2, 0, 1],
    [0, 2, 1, 1],
    [3, 1, 0, 0],
    [2, 3, 0, 0],
    [2, 2, 1, 0],
    [1, 1, 1, 0],
    [4, 3, 0, 0],
    [2, 2, 0, 1]
];

function analise() {
    var verticalLines = parseInt(document.getElementById("verticalLines").value);
    var horizontalLines = parseInt(document.getElementById("horizontalLines").value);
    var diagonal1Lines = parseInt(document.getElementById("diagonal1Lines").value);
    var diagonal2Lines = parseInt(document.getElementById("diagonal2Lines").value);

    var resProbabilities = [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0];
    for (var i = 0; i < categories.length; i++) {
        if (categories[i][0] == verticalLines) resProbabilities[i] ++;
        if (categories[i][1] == horizontalLines) resProbabilities[i] ++;
        if (categories[i][2] == diagonal1Lines) resProbabilities[i] ++;
        if (categories[i][3] == diagonal2Lines) resProbabilities[i] ++;
        resProbabilities[i] /= 4.0;
    }

    var maxProbability = 0.0;
    for (var i in resProbabilities)
        if (maxProbability < resProbabilities[i])
            maxProbability = resProbabilities[i];

    document.getElementById("res").innerHTML = "";
    for (var i in resProbabilities) {
        if (resProbabilities[i] == maxProbability)
            document.getElementById("res").innerHTML += "Ймовірність що це цифра " + i + " = " + maxProbability + "\n";
    }
}
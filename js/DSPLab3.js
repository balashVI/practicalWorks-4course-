var app = angular.module('myApp', [], function ($interpolateProvider) {
    $interpolateProvider.startSymbol('[[');
    $interpolateProvider.endSymbol(']]');
});

app.controller('formController', ['$scope', '$http', formController]);

function formController($scope, $http) {
    $scope.numberOfPoints = 8;
    $scope.inputData = [{
            Re: 4,
            Im: 0
        }, {
            Re: 2,
            Im: 0
        }, {
            Re: 1,
            Im: 0
        }, {
            Re: 4,
            Im: 0
        }, {
            Re: 6,
            Im: 0
        }, {
            Re: 3,
            Im: 0
        }, {
            Re: 5,
            Im: 0
        }, {
            Re: 2,
            Im: 0
        }
    ];
    $scope.mode = 0;
    $scope.resData = [];

    $scope.numberOfPointsChanged = function () {
        $scope.inputData = new Array($scope.numberOfPoints);
        for (var i = 0; i < $scope.numberOfPoints; i++) {
            $scope.inputData[i] = {
                Re: 0,
                Im: 0
            };
        }
    }

    $scope.submitForm = function (isValid) {
        if (isValid) {
            $http.get("/DSP/lab3Calc?mode=" + $scope.mode + "&inputData=" + JSON.stringify($scope.inputData))
                .success(function (data, status, headers, config) {
                    $scope.resData = data;
                }).error(function (data, status, headers, config) {
                    // called asynchronously if an error occurs
                    // or server returns response with an error status.
                });
        }
    }

    $http.get("/DSP/lab3Calc2").success(function (data, status, headers, config) {
        createChart(data.InpSignal1, data.Time, "inp1chart");
        createChart(data.ResSignal1, data.Frequency, "res1chart");
        createChart(data.InpSignal2, data.Time, "inp2chart");
        createChart(data.ResSignal2, data.Frequency, "res2chart");
    });
}

function createChart(dataVector, labelsList, canvasId) {
    var inpData = new Array(dataVector.length);
    for (var i = 0; i < dataVector.length; i++)
        inpData[i] = Math.sqrt(dataVector[i].Re * dataVector[i].Re + dataVector[i].Im * dataVector[i].Im)

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
                data: inpData
        }
    ]
    }, null);
}
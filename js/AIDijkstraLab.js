app = angular.module('myApp', []);
app.config(function ($interpolateProvider) {
    $interpolateProvider.startSymbol('[[');
    $interpolateProvider.endSymbol(']]');
});

app.controller('FormController', ['$scope', '$http', function ($scope, $http) {
    $scope.numberOfVertices = 6;
    $scope.counter = [0, 1, 2, 3, 4, 5]
    $scope.cells = [-1, 7, 9, -1, -1, 14, 7, -1, 10, 14, -1, -1, 9, 10, -1, 11, -1, 2, -1, 15, 11, -1, 6, -1, -1, -1, -1, 6, -1, 9, 14, -1, 2, -1, 9, -1]
    $scope.startPoint = 0;
    $scope.finishPoint = 4;
    $scope.currentSlide = 0;
    $scope.slides = [];

    $scope.numbOfVerticesChanged = function () {
        var n = parseInt($scope.numberOfVertices);
        $scope.counter = new Array(n);
        $scope.cells = new Array(n * n);
        for (var i = 0; i < n; i++) {
            $scope.counter[i] = i;
            for (var j = 0; j < n; j++)
                $scope.cells[i * n + j] = -1;
        }
    }

    $scope.calcIndex = function (i, j) {
        if (i < j)
            return i * $scope.numberOfVertices + j;
        else
            return j * $scope.numberOfVertices + i;
    }

    $scope.sliderChanged = function () {
        document.getElementById("svgContainer").innerHTML = $scope.slides[$scope.currentSlide];
    }

    $scope.submitForm = function (isValid) {
        if (isValid) {
            var requestData = {};
            requestData.NumberOfVertices = $scope.numberOfVertices;
            requestData.EdgesMatrix = $scope.cells;
            requestData.StartPoint = $scope.startPoint;
            requestData.FinishPoint = $scope.finishPoint;
            $http.get("/AI/DijkstraLabCalc?requestData=" + JSON.stringify(requestData))
                .success(function (data, status, headers, config) {
                    $scope.currentSlide = 0
                    $scope.startPoint = data.StartPoint;
                    $scope.finishPoint = data.FinishPoint;
                    $scope.cells = data.EdgesMatrix;
                    $scope.distance = data.Distance;
                    $scope.path = data.Path;
                    $scope.slides = data.Slides;
                    document.getElementById("svgContainer").innerHTML = data.Slides[0];

                }).error(function (data, status, headers, config) {
                    // called asynchronously if an error occurs
                    // or server returns response with an error status.
                });
        }

    };
}]);
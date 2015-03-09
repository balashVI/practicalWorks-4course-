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
}
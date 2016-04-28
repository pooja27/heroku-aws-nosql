/**
 * Created by Jagmohan on 4/9/16.
 */

var app = angular.module('loginApp', []);
app.controller('loginController', function ($http,$scope)
{alert("HEllo");


    $scope.login=function(){
        $http({
            method: 'POST',
            url: '/login',
            data: $scope.userLogin
        }).
        then(function(response) {

          alert("Login Successful");

        },function(response){
            alert("Login Failure");
        });

    }
    $scope.register=function(){
        $http({
            method: 'POST',
            url: '/signup',
            data: $scope.userRegister
        }).
        then(function(response) {

            alert("Registeration Successful");

        },function(response){
            alert("Registeration Failure");
        });

    }
});
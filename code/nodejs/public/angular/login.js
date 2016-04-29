/**
 * Created by Jagmohan on 4/9/16.
 */

var app = angular.module('loginApp', []);
app.controller('loginController', function ($http,$scope)
{

    $http({
        method: 'GET',
        url: '/allProducts',

    }).
    then(function(response) {

        alert("Product Successful");
        alert(response.data);
        console.log(response.data);
        $scope.products=response.data;

    },function(response){
        alert("Failure");
    });




    //$scope.login=function(){
    //    $http({
    //        method: 'POST',
    //        url: '/login',
    //        data: $scope.userLogin
    //    }).
    //    then(function(response) {
    //
    //      alert("Login Successful");
    //        alert(response.data);
    //        console.log(response.data);
    //
    //    },function(response){
    //        alert("Login Failure");
    //        console.log(response.data);
    //    });
    //
    //}
    //$scope.register=function(){
    //    $http({
    //        method: 'POST',
    //        url: '/signup',
    //        data: $scope.userRegister
    //    }).
    //    then(function(response) {
    //
    //        alert("Registeration Successful");
    //        alert(response.data);
    //
    //    },function(response){
    //        alert("Registeration Failure");
    //    });
    //
    //}
});
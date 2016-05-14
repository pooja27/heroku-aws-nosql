// Module for the demo

angular.module('splashDemo', ['ui.splash'])
.controller('MainCtrl', ['$splash','$scope','$http','$window', function ($splash,$scope,$http,$window) {
  this.openSplash = function () {
    $splash.open({
      title: 'Welcome to Starbucks!',
      message: "Please enter details below"
    });
  };
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

  $scope.login=function()
  {
    $http({
      method: 'POST',
      url: '/login',
      data: $scope.userLogin
    }).
    then(function(response) {



      console.log(response.data);
      $window.location.href = '/';
    },function(response){
      alert("Login Failure");
      console.log(response.data);
    });
  }
}]);

// Re-usable $splash module
angular.module('ui.splash', ['ui.bootstrap', 'ngAnimate'])
.service('$splash', [
  '$uibModal',
  '$rootScope',
  function($uibModal, $rootScope) {
    return {
      open: function (attrs, opts) {
        var scope = $rootScope.$new();
        angular.extend(scope, attrs);
        opts = angular.extend(opts || {}, {
          backdrop: false,
          scope: scope,
          templateUrl: 'splash/content.html',
          windowTemplateUrl: 'splash/home1.ejs'
        });
        return $uibModal.open(opts);
      }
    };
  }
])
.run([
  '$templateCache',
  function ($templateCache) {
    $templateCache.put('splash/home1.ejs',
      '<section modal-render="{{$isRendered}}" class="splash" modal-in-class="splash-open" ng-style="{\'z-index\': 1000, display: \'block\'}" ng-click="close($event)">' +
      '  <div class="splash-inner" ng-transclude></div>' +
      '</section>'
    );
    $templateCache.put('splash/content.html',
	'<form id="registrationform" method="POST" name="regform" action="\signup">'+
      '<div class="splash-content text-center">' +
      '<h1 ng-bind="title"></h1>' +
      '<p class="lead" ng-bind="message"></p>' +
	  'First Name:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<input type="text" name="uname" id="uname" style="background-color : #000000; border-radius:5px;" required/><br><br>'+
    'Last Name:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<input type="text" name="lname" id="lname" style="background-color : #000000; border-radius:5px;" required/><br><br>'+
    'Password:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; <input type="password" name="password" id="password" style="background-color : #000000;border-radius:5px;" required/><br><br>'+
	  'Email ID:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<input type="email" id="email" name="email" style="background-color : #000000; border-radius:5px;" required /><br><br>'+
      '<button class="btn btn-lg btn-outline" ng-bind="btnText || \'Register\'" ></button>' +
      '</div>'+
	  '</form>'+
	  '<button class="fa fa-times" style="margin-left: 900px;  margin-top: 120px;" ng-bind="btnText || \'Cancel\'" ng-click="$close()"></button>'
    );
  }
]);


var app = angular.module('myApp', []);
app.controller("Controller", ['$scope','$http', function($scope,$http) {
    $http({
        method: 'GET',
        url: '/alltea',

    }).
    then(function(response) {

        console.log(response.data);
        $scope.teacataloge=response.data;

    },function(response){
        alert("Failure");
    });
    $http({
        method: 'GET',
        url: '/allDrinkware',

    }).
    then(function(response) {

        console.log(response.data);
        $scope.drinkcataloge=response.data;

    },function(response){
        alert("Failure");
    });
    $http({
        method: 'GET',
        url: '/allCoffee',

    }).
    then(function(response) {

        console.log(response.data);
        $scope.coffecataloge=response.data;

    },function(response){
        alert("Failure");
    });
    //$scope.coffecataloge = [{
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Sulawesi, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "3 Region Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Eastern D.R. Congo Lake Kivu",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Reserve Burundi Murambi",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Malawi Sable Farms",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Espresso Roast Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Espresso Roast Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Sumatra K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Caffè Verona K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Ready Brew French Roast",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Pike Place Roast K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Pike Place Roast Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Caffè Verona Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks VIA Ready Brew Italian Roast",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Espresso Roast, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Caffè Verona Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Pike Place Roast Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Ready Brew French Roast",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Caffé Verona Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Caffè Verona K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Milk Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks VIA Ready Brew Decaf Italian Roast",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Guatemala Antigua Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Sumatra Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Pike Place Roast, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks House Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Guatemala Antigua Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks House Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Ready Brew Decaf Italian Roast",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Colombia Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks French Roast K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Milk Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Komodo Dragon Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Ethiopia Yirgacheffe Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks House Blend Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Sumatra, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Veranda Blend Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Colombia Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Torrefazione Italia Napoli Blend, Ground",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks French Roast, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Torrefazione Italia Perugia Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Caffè Latte Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Sumatra Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "House Blend Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Roastery - Ongoing Subscription",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks House Blend, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Veranda Blend",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Veranda Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "House Blend Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Decaf Espresso Roast Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks VIA Caramel Flavored Iced Coffee",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Pike Place Roast, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Refreshers Strawberry Lemonade",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Ethiopia Yirgacheffe Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Gold Coast Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Veranda Blend",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks VIA Latte Vanilla Latte",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Veranda Blend Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Pike Place Roast K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Sumatra Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks VIA Instant Caramel Latte",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Italian Roast, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Sumatra, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks French Roast K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Breakfast Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks VIA Refreshers Valencia Orange",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Decaf Espresso Roast Espresso Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Pike Place Roast Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Caffé Verona Blend, Ground",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Sumatra K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Latte Caffè Mocha",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Roastery - 3 Month Subscription",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks VIA Instant White Chocolate Mocha Latte",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Decaf Pike Place Roast, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Veranda Blend, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Decaf Pike Place Roast Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks French Roast, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks House Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Decaf Caffè Verona,&nbsp;Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Veranda Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Vietnam Da Lat",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Breakfast Blend, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Decaf Pike Place Roast Brewed Coffee Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Kenya, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Sumatra Aceh",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Decaf House Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "French Roast Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Guatemala Antigua, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks VIA Ready Brew Colombia",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Cameroon Mt. Oku",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Decaf Sumatra, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Organic Yukon Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Decaf House Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Verismo Caramel Macchiato Kit",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Ethiopia, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Caffè Verona Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Sumatra K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Decaf Espresso Roast, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Decaf Pike Place Roast Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Caffè Latte Verismo Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Verismo White Chocolate Kit",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Veranda Blend Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Fair Trade Certified Cafè Estima Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Single-Origin Coffee Variety Pack",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Caramel K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Medium Roast Coffee Variety Pack",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Puerto Rico, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Decaf Pike Place Roast, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Veranda Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Torrefazione Italia Montecatini Decaf Blend, Ground",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Roastery - 12 Month Subscription",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Bright Sky Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Reserve Roastery - 6 Month Subscription",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks Pike Place Roast K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Vanilla K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Willow Blend, Whole Bean",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Mocha K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Tazo Awake English Breakfast K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks French Roast K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks House Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Dark Roast Coffee Variety Pack",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Fair Trade Certified  Cafè Estima Blend Portion Packs",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks K-Cup Pods Sampler - French Roast &amp; Sumatra",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Starbucks K-Cup Pods Sampler - Pike Place Roast &amp; Caffè Verona",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Verismo Vanilla Latte Kit",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Caramel K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Reserve Decaf Costa Rica Bella Vista F.W. Tres Rios",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Caffè Verona K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "asia-pacific",
    //    "category": "coffee",
    //    "price": 20,
    //    "name": "Verismo Mocha Kit",
    //    "flavor": "flavored",
    //    "quantity": 3,
    //    "roast": "dark",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Sumatra K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks K-Cup Pods Sampler - French Roast &amp; Caffè Verona",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks K-Cup Pods Sampler - French Roast &amp; Pike Place Roast",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks K-Cup Pods Sampler - House Blend &amp; Pike Place Roast",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "latin america",
    //    "category": "coffee",
    //    "price": 40,
    //    "name": "Starbucks Vanilla K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "blonde",
    //    "type": "decaffinated"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Mocha K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}, {
    //    "region": "multi",
    //    "category": "coffee",
    //    "price": 50,
    //    "name": "Starbucks Bright Sky Blend K-Cup Pods",
    //    "flavor": "flavored",
    //    "quantity": 5,
    //    "roast": "medium",
    //    "type": "regular"
    //}];
}]);






/* ========================================================================= */
/*	Preloader
/* ========================================================================= */

jQuery(window).load(function(){

	$("#preloader").fadeOut("slow");

});

/* ========================================================================= */
/*  Welcome Section Slider
/* ========================================================================= */

$(function() {

    var Page = (function() {

        var $navArrows = $( '#nav-arrows' ),
            $nav = $( '#nav-dots > span' ),
            slitslider = $( '#slider' ).slitslider( {
                onBeforeChange : function( slide, pos ) {

                    $nav.removeClass( 'nav-dot-current' );
                    $nav.eq( pos ).addClass( 'nav-dot-current' );

                }
            } ),

            init = function() {

                initEvents();
                
            },
            initEvents = function() {

                // add navigation events
                $navArrows.children( ':last' ).on( 'click', function() {

                    slitslider.next();
                    return false;

                } );

                $navArrows.children( ':first' ).on( 'click', function() {
                    
                    slitslider.previous();
                    return false;

                } );

                $nav.each( function( i ) {
                
                    $( this ).on( 'click', function( event ) {
                        
                        var $dot = $( this );
                        
                        if( !slitslider.isActive() ) {

                            $nav.removeClass( 'nav-dot-current' );
                            $dot.addClass( 'nav-dot-current' );
                        
                        }
                        
                        slitslider.jump( i + 1 );
                        return false;
                    
                    } );
                    
                } );

            };

            return { init : init };

    })();

    Page.init();

});



$(document).ready(function(){

	/* ========================================================================= */
	/*	Menu item highlighting
	/* ========================================================================= */

	jQuery('#nav').singlePageNav({
		offset: jQuery('#nav').outerHeight(),
		filter: ':not(.external)',
		speed: 2000,
		currentClass: 'current',
		easing: 'easeInOutExpo',
		updateHash: true,
		beforeStart: function() {
			console.log('begin scrolling');
		},
		onComplete: function() {
			console.log('done scrolling');
		}
	});
	
    $(window).scroll(function () {
        if ($(window).scrollTop() > 400) {
            $(".navbar-brand a").css("color","#fff");
            $("#navigation").removeClass("animated-header");
        } else {
            $(".navbar-brand a").css("color","inherit");
            $("#navigation").addClass("animated-header");
        }
    });
	
	/* ========================================================================= */
	/*	Fix Slider Height
	/* ========================================================================= */	

    // Slider Height
    var slideHeight = $(window).height();
    
    $('#home-slider, #slider, .sl-slider, .sl-content-wrapper').css('height',slideHeight);

    $(window).resize(function(){'use strict',
        $('#home-slider, #slider, .sl-slider, .sl-content-wrapper').css('height',slideHeight);
    });
	
	
	
	$("#works, #testimonial").owlCarousel({	 
		navigation : true,
		pagination : false,
		slideSpeed : 700,
		paginationSpeed : 400,
		singleItem:true,
		navigationText: ["<i class='fa fa-angle-left fa-lg'></i>","<i class='fa fa-angle-right fa-lg'></i>"]
	});
	
	
	/* ========================================================================= */
	/*	Featured Project Lightbox
	/* ========================================================================= */

	$(".fancybox").fancybox({
		padding: 0,

		openEffect : 'elastic',
		openSpeed  : 650,

		closeEffect : 'elastic',
		closeSpeed  : 550,

		closeClick : true,
			
		beforeShow: function () {
			this.title = $(this.element).attr('title');
			this.title = '<h3>' + this.title + '</h3>' + '<p>' + $(this.element).parents('.portfolio-item').find('img').attr('alt') + '</p>';
		},
		
		helpers : {
			title : { 
				type: 'inside' 
			},
			overlay : {
				css : {
					'background' : 'rgba(0,0,0,0.8)'
				}
			}
		}
	});
	
});


/* ==========  START GOOGLE MAP ========== */

// When the window has finished loading create our google map below
google.maps.event.addDomListener(window, 'load', init);

function init() {
    // Basic options for a simple Google Map
    // For more options see: https://developers.google.com/maps/documentation/javascript/reference#MapOptions

	    var myLatLng = new google.maps.LatLng(22.402789, 91.822156);

	    var mapOptions = {
	        zoom: 15,
	        center: myLatLng,
	        disableDefaultUI: true,
	        scrollwheel: false,
	        navigationControl: true,
	        mapTypeControl: false,
	        scaleControl: false,
	        draggable: true,

        // How you would like to style the map. 
        // This is where you would paste any style found on Snazzy Maps.
        styles: [{
            featureType: 'water',
            stylers: [{
                color: '#46bcec'
            }, {
                visibility: 'on'
            }]
        }, {
            featureType: 'landscape',
            stylers: [{
                color: '#f2f2f2'
            }]
        }, {
            featureType: 'road',
            stylers: [{
                saturation: -100
            }, {
                lightness: 45
            }]
        }, {
            featureType: 'road.highway',
            stylers: [{
                visibility: 'simplified'
            }]
        }, {
            featureType: 'road.arterial',
            elementType: 'labels.icon',
            stylers: [{
                visibility: 'off'
            }]
        }, {
            featureType: 'administrative',
            elementType: 'labels.text.fill',
            stylers: [{
                color: '#444444'
            }]
        }, {
            featureType: 'transit',
            stylers: [{
                visibility: 'off'
            }]
        }, {
            featureType: 'poi',
            stylers: [{
                visibility: 'off'
            }]
        }]
    };

    // Get the HTML DOM element that will contain your map 
    // We are using a div with id="map" seen below in the <body>
    var mapElement = document.getElementById('map-canvas');

    // Create the Google Map using our element and options defined above
    var map = new google.maps.Map(mapElement, mapOptions);

    // Let's also add a marker while we're at it
    var marker = new google.maps.Marker({
        position: new google.maps.LatLng(22.402789, 91.822156),
        map: map,
		icon: 'img/icons/map-marker.png',
    });
}

// ========== END GOOGLE MAP ========== //

var wow = new WOW ({
	offset:       75,          // distance to the element when triggering the animation (default is 0)
	mobile:       false,       // trigger animations on mobile devices (default is true)
});
wow.init();


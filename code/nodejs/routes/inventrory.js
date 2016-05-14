/**
 * Created by Jagmohan on 4/9/16.
 */

var http=require("http");
var Mongo=require('./mongo');
exports.teas=function(req,res){


    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/teas",
        method: 'GET'
    };

    callback = function(response) {
        var str = '';

        console.log(response.statusCode);
        response.on('error',function(){
            console.log("Error in response: "+"\n"+str);

        })
        response.on('data', function (chunk) {
            str += chunk;
        });


        response.on('end', function () {
            var data = JSON.parse(str);

                res.status(200).send(data);

        });
    }

    http.get(options, callback).end();

};
exports.coffees=function(req,res){


    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/coffees",
        method: 'GET'
    };

    callback = function(response) {
        var str = '';

        console.log(response.statusCode);
        response.on('error',function(){
            console.log("Error in response: "+"\n"+str);

        })
        response.on('data', function (chunk) {
            str += chunk;
        });


        response.on('end', function () {
            var data = JSON.parse(str);

            res.status(200).send(data);

        });
    }

    http.get(options, callback).end();

};
exports.drinkwares=function(req,res){


    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/drinkwares",
        method: 'GET'
    };

    callback = function(response) {
        var str = '';

        console.log(response.statusCode);
        response.on('error',function(){
            console.log("Error in response: "+"\n"+str);

        })
        response.on('data', function (chunk) {
            str += chunk;
        });


        response.on('end', function () {
            var data = JSON.parse(str);

            res.status(200).send(data);

        });
    }

    http.get(options, callback).end();

};




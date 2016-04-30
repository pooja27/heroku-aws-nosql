/**
 * Created by Jagmohan on 4/9/16.
 */

var http=require("http");
exports.teas=function(req,res){


    var options = {
        host: 'ec2-52-202-168-18.compute-1.amazonaws.com',
        port: 7777,
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
        host: 'ec2-52-202-168-18.compute-1.amazonaws.com',
        port: 7777,
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
        host: 'ec2-52-202-168-18.compute-1.amazonaws.com',
        port: 7777,
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




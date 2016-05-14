/**
 * Created by Jagmohan on 4/9/16.
 */
var http= require("http");
var Mongo=require('./mongo');
var async=require('async');
exports.addCart=function(req,res){
    var item = req.param('item_id');
    var email = req.session.data.email;
    var category=req.param('category');
    var quan=req.param('quantity');
    var count=0;

    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/cart/addItem/"+email+"/"+category+"/"+item,
        method: 'PUT'
    };

    async.whilst(
        function () { return count < quan; },
        function (callback) {
            count++;
            console.log(count);

            http.request(options,function(response){
                var str='';

                response.on('error',function(){
                    console.log("Error in response: "+"\n"+str);

                })
                response.on('data', function (chunk) {
                    str += chunk;
                });


                response.on('end', function () {

                    var data = JSON.parse(str);

                    console.log(response.statusCode);
                    callback(null,count,response,data);


                });

            }).end();
        },
        function (err, n,response,data) {


            if(response.statusCode===200)
                res.status(200).send(data);
            else
                res.status(404).send({"data":"Failed to Get Cart"});
        }
    );


};
exports.getCart=function(req,res){

    var email = req.session.data.email;
console.log(email);
    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/cart/"+email,
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

            console.log(response.statusCode);
            if(response.statusCode===200)
            { console.log(data.items);
            console.log(data.items.length);
                res.status(200).send(data);}
            else
                res.status(404).send({"data":"Failed to Get Cart"});
        });
    }

    http.get(options, callback).end();

};
exports.removeCart=function(req,res){
    var item = req.param('item_id');
    var email = req.session.data.email;
    var quan=req.param('quantity');
    var count=0;

    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/cart/removeItem/"+email+"/"+item,
        method: 'PUT'
    };

    async.whilst(
        function () { return count < quan; },
        function (callback) {
            count++;
            console.log(count);

            http.request(options,function(response){
                var str='';

                response.on('error',function(){
                    console.log("Error in response: "+"\n"+str);

                })
                response.on('data', function (chunk) {
                    str += chunk;
                });


                response.on('end', function () {

                    var data = JSON.parse(str);

                    console.log(response.statusCode);
                    callback(null,count,response,data);


                });

            }).end();
        },
        function (err, n,response,data) {


            if(response.statusCode===200)
                res.status(200).send(data);
            else
                res.status(404).send({"data":"Failed to Get Cart"});
        }
    );


};
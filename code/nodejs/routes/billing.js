/**
 * Created by Jagmohan on 4/9/16.
 */
var http= require("http");
var Mongo=require('./mongo');
var async=require('async');
exports.checkOut=function(req,res){

    var email = req.session.data.email;
    console.log(email);
    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/cart/"+email,
        method: 'DELETE'
    };

    httpGet = function(response) {
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
            {
                console.log("Success");}
            else
                console.log("Fail");
        });
    }
    http.get(options, httpGet).end();

    var status=[];
    async.each(req.body,function(items,callback){
    var count=0;
        var stat=true;

        console.log(items.count);
        console.log(items.item_id);

        async.whilst(
            function () { return count < items.count; },
            function (callback) {
                count++;
                console.log(count);
                var options1 = {
                    host: 'ec2-52-202-168-18.compute-1.amazonaws.com',
                    port: Mongo.PORT,
                    path: "/mongoserver/inventory/removeOne/"+items.item_id,
                    method: 'PUT'
                };

                http.request(options1,function(response){
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
                    res.status(200).send({"data":"Your Order is successful"});
                else
                {
                   status.push({"dataa":"dskjfdskkj"});
                    res.status(404).send({"data":"Your Order might be delayed"});
                }
            }
        );

callback(null,status);
    },function(err,status){
        if(err){
            console.log("Error Occured");
        }
        else{console.log("jo");
            //if(status.length>0)
                //res.status(404).send({"data":"Your Order might be delayed"});

           // else
                //res.status(200).send({"data":"Order Successful"});

        }

    });






};

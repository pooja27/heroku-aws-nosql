/**
 * Created by Jagmohan on 4/9/16.
 */
var mongo = require("http");
exports.addCart=function(req,res){
    var item = req.param('item_id');
console.log(item);
    res.status(200).send({"data":"Success"});
    //
    //var options = {
    //    host: 'ec2-52-202-168-18.compute-1.amazonaws.com',
    //    port: 7777,
    //    path: "/mongoserver/login/"+email,
    //    method: 'GET'
    //};
    //
    //callback = function(response) {
    //    var str = '';
    //
    //    console.log(response.statusCode);
    //    response.on('error',function(){
    //        console.log("Error in response: "+"\n"+str);
    //
    //    })
    //    response.on('data', function (chunk) {
    //        str += chunk;
    //    });
    //
    //
    //    response.on('end', function () {
    //        var data = JSON.parse(str);
    //        if(password===data.password)
    //        {req.session.data=data;
    //            res.status(200).send(data);}
    //        else
    //            res.status(404).send({"data":"Incorrect Password"});
    //    });
    //}
    //
    //http.get(options, callback).end();

};

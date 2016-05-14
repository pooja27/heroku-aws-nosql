/**
 * Created by Jagmohan on 4/9/16.
 */
var http=require("http");
var Mongo= require('./mongo');
exports.login1=function(req,res){
    var email = req.param('username');
    req.session.data=email;
    res.status(200).send({"data":email});
};
exports.login=function(req,res){
    var email = req.param('username');
    var password = req.param('password');

    var options = {
        host: Mongo.URL,
        port: Mongo.PORT,
        path: "/mongoserver/login/"+email,
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
       if(password===data.password)
       {req.session.data=data;
           res.status(200).send(data);}
        else
            res.status(404).send({"data":"Incorrect Password"});
    });
}

http.get(options, callback).end();

};

exports.signup=function(req,res) {
   var email = req.param('email');
   var password = req.param('password');
   var firstName = req.param('uname');
   var lastName = req.param('lname');

var query=JSON.stringify({
    "userid" : email,
    "password" : password,
    "email" : email,
    "name" : firstName+" "+lastName
    }
);

var options = {
    host: Mongo.URL,
    port: Mongo.PORT,
    path: '/mongoserver/signup',
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Content-Length': query.length
    }

};
var reqPost = http.request(options, function (response) {

    if(response.statusCode===201)
    res.render('login');
    else
    res.render('login');
    response.on('data', function (data) {
        console.log('Posting Result:\n');
        process.stdout.write(data);
        console.log('\n\nPOST Operation Completed');
    });
});

// 7
reqPost.write(query);
reqPost.end();
reqPost.on('error', function (e) {
    console.error(e);
});
};
exports.logout=function(req,res){
    req.session.destroy();
    res.render('login');

}



/**
 * Created by Jagmohan on 4/9/16.
 */
var http=require("http");
exports.login=function(req,res){
    var email = req.param('username');
    var password = req.param('password');

    var options = {
        host: 'ec2-52-202-168-18.compute-1.amazonaws.com',
        port: 7777,
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
   var email = req.param('username');
   var password = req.param('password');
   var firstName = req.param('firstName');
   var lastName = req.param('lastName');

var query=JSON.stringify({
    "userid" : email,
    "password" : password,
    "email" : email,
    "name" : firstName+" "+lastName
    }
);

var options = {
    host: 'ec2-52-202-168-18.compute-1.amazonaws.com',
    port: 7777,
    path: '/mongoserver/signup',
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Content-Length': query.length
    }

};
var reqPost = http.request(options, function (response) {

    if(response.statusCode===201)
    res.status(200).send({"data":"Registration Successful"});
    else
    res.status(404).send({"data":"Registration Failed"});
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



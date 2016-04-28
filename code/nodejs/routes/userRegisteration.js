/**
 * Created by Jagmohan on 4/9/16.
 */
var http=require("http");
exports.login=function(req,res){
    var email = req.param('username');
    var password = req.param('password');
    //var query=JSON.stringify({
    //        "userid" : "jagg",
    //
    //    }
    //);
    //var postheaders = {
    //    'Content-Type' : 'application/json',
    //    'Content-Length' : Buffer.byteLength(query)
    //};
    var options = {
        host: "ec2-54-175-24-14.compute-1.amazonaws.com",
        port: 7777,
        path: "/mongoserver/login/"+"jagg@sjsu.edu",
        method: 'GET'
    };

callback = function(response) {
    var str = '';

// str+="statusCode:"+response.statusCode+"\n";
    console.log(response.statusCode);
    response.on('error',function(){
        console.log("Error in response: "+"\n"+str);

    })
    response.on('data', function (chunk) {
        str += chunk+"\n";
    });


    response.on('end', function () {
        console.log(str);
    });
}

http.get(options, callback).end();
    //mongo.connect(mongoURL, function(){
    //    console.log('Connected to mongo at: ' + mongoURL);
    //    var coll = mongo.collection('userDetails');
    //    coll.findOne({"email:": email, "password":password }, function(err, user){
    //        if(user)
    //        {
    //            console.log(user.email);
    //            res.status(200).send({"status":"Login Successful"});
    //        }
    //        else
    //        {
    //             res.status(401).send({"status":"Login Failed"});
    //
    //        }
    //
    //    });
    //
    //});
    //

};

exports.signup=function(req,res) {
//    var email = req.param('username');
//    var password = req.param('password');
//    var firstName = req.param('firstName');
//    var lastName = req.param('lastName');
//    var mobileNumber = req.param('mobileNumber');
//    console.log(email);
//    console.log(password);
//    console.log(firstName);
//    console.log(lastName);
//    console.log(mobileNumber);
var query=JSON.stringify({
        "userid" : "poo",
        "password":"poom@sjsu.edu",
        "name" : "Pooja"
    }
);

var options = {
    host: 'ec2-54-175-24-14.compute-1.amazonaws.com',
    port: 7777,
    path: '/mongoserver/signup',
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Content-Length': query.length
    }

};
var reqPost = http.request(options, function (res) {
    console.log("response statusCode: ", res.statusCode);
    res.on('data', function (data) {
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



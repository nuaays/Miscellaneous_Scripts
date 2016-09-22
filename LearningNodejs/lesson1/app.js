var express = require('express');
var app = express();

app.get("/", function(req, res){
   res.send('Hello World');
   //res.send(req.headers);

});


app.listen(3000, function(){
   console.log('app is listening at port 3000.');
});

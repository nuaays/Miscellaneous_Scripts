//event driven

mongoose.connect('mongodb://127.0.0.1:27017/test');
mongoose.connection.on('open', function(){
	console.log("Connection to Mongoose ...");
});
//asynchronous


//require and modules

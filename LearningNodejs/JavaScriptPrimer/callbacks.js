//
//Anonymous functions and callbacks


//timeout is 5000ms
console.log("Hello ... ...");
setTimeout(function() {
	console.log("World!");
}, 5000);

//change value after some seconds
var someValue = 100;
var myFunction=function(callback){
	setTimeout(function(){
		someValue = someValue / 2;
	}, 5000);
}

myFunction(function(){
	console.log(someValue);
});

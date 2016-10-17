var worker = require('./dowork');

var something = 1;
var somethingElse = 2;

var newVal = worker.dowork(something, somethingElse);
console.log(newVal);
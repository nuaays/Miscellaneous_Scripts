var myObject = {};


function fullName(firstName, lastName) {
	return  firstName + ' ' + lastName
}


var person = {};
person.firstName = "Jason";
person.lastName = "Krol";
person.fullName = function() {
	return this.firstName + ' ' + this.lastName;
}
person.colors = ['red', 'blue', 'green'];


//
var book = {
	titile: "Web Development with MongoDB and NodeJS",
	author: "Jason Krol",
	publisher: "Packet Publishing"
}
console.log(book.titile);
book.pageCount = 150;
console.log(book);


//nested object
var price = {
	value: '100 RMB'
}
var book = {
	titile: "Web Development with MongoDB and NodeJS",
	author: "Jason Krol",
	publisher: "Packet Publishing",
	charge: price
}
console.log(book.charge.value)

//functions are objects, and private property
var myFunction = function() {
	if(this.timesRun)
		this.timesRun += 1
	else
		this.timesRun = 1
	console.log(timesRun);
}
myFunction();
myFunction();
myFunction();
console.log(myFunction.timesRun);


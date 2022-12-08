const fs = require('fs');

const contents = fs.readFileSync('./input.txt', 'utf8');
var total = 0;
const max = [];
contents.split(/\r?\n/).forEach(line => {
	if(line == ''){
		max.push(total);
		total = 0;
	}else{
		line = parseInt(line)
		total += line;
	}
});
max.sort((a,b) => {return a - b})
var one = max.pop();
var two = max.pop();
var three = max.pop();
total = one + two + three;
console.log(max);
console.log("one: ", one);
console.log("two: ", two);
console.log("three: ", three);
console.log("total: ", total);

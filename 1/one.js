const fs = require('fs');

const contents = fs.readFileSync('./input.txt', 'utf8');
var total = 0;
var max = 0;
contents.split(/\r?\n/).forEach(line => {
	if(line == ''){
		if (total > max){
			max = total
		}
		total = 0;
	}else{
		line = parseInt(line)
		total += line;
	}
});
console.log(max)

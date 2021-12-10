const numbers = require("./arr.json");

let total = 0;

for (const i in numbers) {
  total = total + numbers[i];
}

console.log(total);

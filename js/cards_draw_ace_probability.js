let numSims = 10000000000;
let count = 0;

for (let i = 0; i < numSims; i++) {
  let hand = [];
  for (let j = 0; j < 5; j++) {
    let card = Math.floor(Math.random() * 13) + 1;
    hand.push(card);
  }
  if (hand.filter((card) => card === 1).length >= 2) {
    count++;
  }
}

let prob = count / numSims;
console.log(`Probability of drawing at least two aces in a five-card hand 
with replacement: ${prob}`);

const array = [12, 1, 405, 390, 23523, 1093, 3435, 234, 2344, 2345, 3465, 7646];

const main = (x, biggerThan) => {
  const newArray = [];

  for (const i in array) {
    if (array[i] > biggerThan) {
      newArray.push(array[i]);
    }
  }

  console.log("ORIGINAL", array);

  console.log("BIGGER THAN", newArray);

  console.log(
    `ASKED LAST ${x} VALUES`,
    newArray.slice(newArray.length - x, newArray.length)
  );
};

main(4, 300);

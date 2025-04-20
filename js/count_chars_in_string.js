const main = (chars) => {
  const seen = {};

  const array = Array.from(chars);

  for (const i in array) {
    const exists = Object.keys(seen).includes(array[i]);

    if (exists) {
      seen[array[i]] += 1;
    } else {
      seen[array[i]] = 1;
    }
  }

  for (const i in seen) {
    if (seen[i] > 1) {
      console.log(i, seen[i]);
    }
  }
};

main("Hello my name is Viljami, this is a test string");

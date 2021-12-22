const arr1 = [0, 2, 3, 4, 5, 6, 7, 8, 9];
const arr2 = [1, 2, 3, 4, 5, 6, 7, 8, 9];

const main = (arr1, arr2) => {
  const found = arr2.find((value) => arr1[0] === value);

  if (found) {
  } else {
    console.log("First number not found, can't be rotated version");
  }
};

main(arr1, arr2);

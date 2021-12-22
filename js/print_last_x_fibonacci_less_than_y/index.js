const fibo = (x) => {
  if (x <= 1) {
    return x;
  }

  return fibo(x - 1) + fibo(x - 2);
};

const main = (x, y) => {
  const smaller_than = fibo(y);

  console.log(smaller_than);
};

main(2, 3);

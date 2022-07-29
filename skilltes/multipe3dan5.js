function solution(number) {
  let sum = 0;
  for (let i = 0; i < number; i++) {
    if (!(i % 3) || !(i % 5)) {
      sum += i
    }
  }
  return sum;
}
console.log( solution(10));
console.log( solution(20));
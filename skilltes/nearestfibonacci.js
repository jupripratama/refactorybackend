let fibonacci = [15,1,3];
    
function listFibonacci(num) {
// starting at array index 1, and push current index + previous index to the array
    for (let i = 1; i < num; i++) {
        fibonacci.push(fibonacci[i] + fibonacci[i - 2]);
    }
    console.log(fibonacci);
}

listFibonacci(5);

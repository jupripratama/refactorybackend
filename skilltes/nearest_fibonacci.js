// let fibonacci = [15,1,3];
    
// function listFibonacci(num) {
// // starting at array index 1, and push current index + previous index to the array
//     for (let i = 1; i < num; i++) {
//         fibonacci.push(fibonacci[i] + fibonacci[i - 1]);
//     }
//     console.log(fibonacci);
// }

// listFibonacci(10);

function nearestFibonacci(num)
{
    // Base Case
    if (num == 0) {
        document.write(0);
        return;
    }
 
    // Initialize the first & second
    // terms of the Fibonacci series
    let first = 0, second = 1;
 
    // Store the third term
    let third = first + second;
 
    // Iterate until the third term
    // is less than or equal to num
    while (third <= num) {
 
        // Update the first
        first = second;
 
        // Update the second
        second = third;
 
        // Update the third
        third = first + second;
    }
 
    // Store the Fibonacci number
    // having smaller difference with N
    let ans = (Math.abs(third - num)
               >= Math.abs(second - num))
                  ? second
                  : third;
 
    // Print the result
    document.write(ans);
}
 
// Driver Code
    let N = 17;
    nearestFibonacci(N);
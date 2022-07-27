#include <bits/stdc++.h>
using namespace std;

// Function to find the Fibonacci
// number which is nearest to N
void nearestFibonacci(int num)
{
    // Base Case
    if (num == 0)
    {
        cout << 0;
        return;
    }

    // Initialize the first & second
    // terms of the Fibonacci series
    int first = 0, second = 1;

    // Store the third term
    int third = first + second;

    // Iterate until the third term
    // is less than or equal to num
    while (third <= num)
    {

        // Update the first
        first = second;

        // Update the second
        second = third;

        // Update the third
        third = first + second;
    }

    // Store the Fibonacci number
    // having smaller difference with N
    int ans = (abs(third - num) >= abs(second - num))
                  ? second
                  : third;

    // Print the result
    cout << ans;
}

// Driver Code
int main()
{
    int N = 19;
    nearestFibonacci(N);

    return 0;
}
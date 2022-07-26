#include<iostream>
#include<algorithm>
#include <vector>

using namespace std;

int main() {
    int num;
    int dec;
    bool flag = false;
    int curNum;
    vector<int> numbers;
    
    cout << "Enter three digit number:  ";
    cin >> num;

    while (num > 0) {
        curNum= num % 10;
        num /= 10;
        flag = false;
        for (int i = 0; i < chisla.size(); ++i) {
            if (numbers[i] > curNum) {
                flag = true;
                numbers.insert(numbers.begin() + i, curNum);
                break;
            }
        }
        if (!flag) numbers.push_back(curNum);
    }

    // Сортиране на масива, като се подрежда в низходящ ред.
    sort(numbers.begin(), numbers.end(), greater<int>());

   
    dec = 1;
    int min_val = 0;
    for (auto &it : numbers) {
        min_val += it * dec;
        dec *= 10;
    }

    sort(numbers.begin(), numbers.end());

    dec = 1;
    int max_val = 0;
    for (auto &it : numbers) {
        max_val += it * dec;
        dec *= 10;
    }

    cout << "е: \n"
         << (max_val - min_val) << endl;

    return 0;
}
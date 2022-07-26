function productOfArrayExceptSelf(array){
    return array.map(function (_, i) {
        return array.reduce(function (product, val, j) {
            return product * (i === j ? 1 : val);
        }, 1);
    });
}

var array = [12,20];
console.log(productOfArrayExceptSelf(array));
// Sample data
var array = [12,20];
console.log(productOfArrayExceptSelf(array));
// Sample data
var array = [3,27,4,2];
console.log(productOfArrayExceptSelf(array));
// Sample data
var array = [13,10,5,2,9];
console.log(productOfArrayExceptSelf(array));
// Sample data
var array = [16,17,4,3,5,2];
console.log(productOfArrayExceptSelf(array));
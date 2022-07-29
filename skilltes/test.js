var productExceptSelf = function(nums) {
  let leftProduct = [];
  let rightProduct = [];
  let solution = [];
  
  for(let i = 0; i < nums.length; i++){
      if(leftProduct.length===0){
         leftProduct.push(1);
         } else {
             leftProduct.push(leftProduct[i-1]*nums[i-1]);
         }
  }

  for(let i = nums.length-1; i>-1; i--){
      if(rightProduct.length===0){
         rightProduct.push(1);
         } else {
             rightProduct.unshift(rightProduct[0]*nums[i+1]);
         }
  }
  
  for(let i = 0; i < leftProduct.length; i++){
      solution.push(leftProduct[i]*rightProduct[i]);
  }
  
  return solution;

  
};
var nums = [12,20];
console.log(productExceptSelf(nums))
var nums = [3,27,4,2];
console.log(productExceptSelf(nums))
var nums = [3,10,5,2,9];
console.log(productExceptSelf(nums))
var nums = [16,17,4,3,5,2];
console.log(productExceptSelf(nums))


/*

Input: nums = [0,1,2,2,3,0,4,2,3] val = 3
   i
* [0,1,2,2,3,0,4,2] iter 1
   j

             i!
* [0,1,2,2,3,0,4,2] iter 1
         j

* */


/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
function removeElementTwoPointers(nums, val) {
  if (nums.length === 0) {
    return 0
  }

  let i = 0
  let j = 0
  for (; i < nums.length; i++) {
    console.log("nums[i]", nums[i])
    if (nums[i] !== val) {
      nums[j++] = nums[i]
    }
  }
  return j
}


/*
53ms
* var removeElement = function(nums, val) {
    let i = 0; // Pointer for non-'val' elements
    for (let j = 0; j < nums.length; j++) {
        if (nums[j] !== val) {
            nums[i] = nums[j];
            i++;
        }
    }
    return i;
};

* */


module.exports = { removeElementTwoPointers}
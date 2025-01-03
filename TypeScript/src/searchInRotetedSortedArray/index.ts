

/* 
    [4,5,6,7,0,1,2]
    target = 6
    left = 4
    middle = 7
    4(left) <= 7(middle)


*/

export const searchInRotatedSortedArray = (nums: number[], target: number): number => {
    if (nums.length === 0) {
        return -1
    }
    
    let left = 0
    let right = nums.length - 1

    while(left <= right) {
        let middle = Math.floor((left + right) / 2)
        if (nums[middle] === target) {
            return middle
        }

        if (nums[left] <= nums[middle]) {
            if(nums[left] <= target && nums[middle] > target) {
                 right = middle - 1
            } else {
                left = middle + 1
            } 
        } else {
            if (nums[middle] < target && target <= nums[right]) {
                left = middle + 1
            } else {
                right = middle - 1
            }
        }
    }

    return -1
} 
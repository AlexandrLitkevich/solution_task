

/*
   i
* [0,0,1,1,1,2,2,3,3,4] case
       j
     i
* [0,1,1,1,1,2,2,3,3,4] case
         j
       i
* [0,1,2,1,1,2,2,3,3,4] case
               j
           i
* [0,1,2,3,4,2,2,3,3,4] case
                     j

* */

//https://leetcode.com/tag/two-pointers/


export function removeDuplicates (nums: number[]): number {
  if (nums?.length === 0) {
    return 0
  }

  let  i = 0
  let j =  1
  const len = nums?.length

  while (j < len) {
      if (nums[i] !==  nums[j]) {
        i++
        nums[i] = nums[j]
      }
      j++
  }
  return i + 1
}

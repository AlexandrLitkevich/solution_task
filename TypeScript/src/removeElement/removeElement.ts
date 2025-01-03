
export function removeElementTwoPointers(nums:number[], val: number): number {
  if (nums.length === 0) return 0

  let i = 0
  let j = 0

  for (; i < nums.length; i++) {
    if (nums[i] !== val) {
      nums[j++] = nums[i]
    }
  }
  return j
}

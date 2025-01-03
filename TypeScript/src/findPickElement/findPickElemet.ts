

export const findPickElement = (nums: number[]): number => {
    if (nums.length === 0) {
        return -1
    }

    let l = 0
    let r = nums.length - 1
    while(l < r) {
        let mid = l + Math.floor((r - l) / 2)

        if (nums[mid] > nums[mid + 1]) {
            r = mid
        } else if (nums[mid] <= nums[mid + 1]) {
            l = mid + 1
        }
    }

    return l
}
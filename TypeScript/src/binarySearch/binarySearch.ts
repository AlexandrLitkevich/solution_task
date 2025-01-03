

export const binarySearch = (arr: number[], target: number): number => {
    let left = 0
    let right = arr.length - 1

    while (left <= right) {
        const middle = Math.floor((left + right) / 2)

        if(arr[middle] === target) {
            return middle
        }

        if(arr[middle] > target) {
            right = middle - 1
        }

        if (arr[middle] < target) {
            left = middle + 1
        }
    }

    return -1
}
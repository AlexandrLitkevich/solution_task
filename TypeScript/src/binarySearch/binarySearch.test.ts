import { binarySearch } from './binarySearch'


describe("binary search", () => {
    test("empty array", () => {
        expect(binarySearch([], 4)).toBe(-1)
    })
    test("success", () => {
        expect(binarySearch([1,2,3,4,5,6,7,8,9,10], 4)).toBe(3)
    })
})


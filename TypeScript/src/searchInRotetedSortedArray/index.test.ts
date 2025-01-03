import { searchInRotatedSortedArray } from '.'


describe("searchInRotatedSortedArray", () => {
    test("empty array", () => {
        expect(searchInRotatedSortedArray([], 4)).toBe(-1)
    })
    test("not target", () => {
        expect(searchInRotatedSortedArray([4,5,6,7,0,1,2], 3)).toBe(-1)
    })
    test("once element", () => {
        expect(searchInRotatedSortedArray([1], 4)).toBe(-1)
    })
    test("success first", () => {
        expect(searchInRotatedSortedArray([4,5,6,7,0,1,2], 6)).toBe(2)
    })
    test("success second", () => {
        expect(searchInRotatedSortedArray([5,1,3], 5)).toBe(0)
    })
})

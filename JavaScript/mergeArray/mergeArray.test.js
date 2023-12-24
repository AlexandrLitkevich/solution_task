const {mergeArrays} = require('./mergeArrays')
const {describe, test, expect} = require("@jest/globals");


describe("merge sort test", () => {
    test("two array", () => {

        const result = mergeArrays([1,2,3,4], [5,6,7,8])
        expect(result.length).toBe(8)
    })
})
const {sum} = require('./sum')
const {describe, test, expect} = require("@jest/globals");

describe("first test", () => {
    test("1 + 1", () => {
        expect(sum(1,1)).toBe(2)
    })
})




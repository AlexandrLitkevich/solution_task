
const {describe, test, expect} = require("@jest/globals");
const {convertHashToArray} = require("./convertHashToArray");

describe("first test", () => {
    test("hash", () => {
        const hash = { "Alex": 23 }
        const expected = [["Alex", 23]]
        expect(convertHashToArray(hash)).toEqual(expect.arrayContaining(expected))
    })
    test("hash nullble", () => {
        const hash = null
        const expected = []
        expect(convertHashToArray(hash)).toEqual(expect.arrayContaining(expected))
    })
    test("big hash", () => {
        const hash = { "Alex": 23, "Vike": 33, "Olga": 55 }
        const expected = [["Alex", 23], ["Vike", 33],["Olga", 55],]
        expect(convertHashToArray(hash)).toEqual(expect.arrayContaining(expected))
    })
})

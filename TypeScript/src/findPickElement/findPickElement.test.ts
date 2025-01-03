
import { findPickElement } from './findPickElemet'

describe("findPickElememnt", () => {
    test("empty array", () => {
        expect(findPickElement([])).toBe(-1)
    })
    test("success", () => {
        expect(findPickElement([1,2,3,1])).toBe(2)
    })

})
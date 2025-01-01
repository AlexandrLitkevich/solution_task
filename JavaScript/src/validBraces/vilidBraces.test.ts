import { validBraces } from './validBraces'



describe("testing validBraces", () => {
    test("{}[]()", () => {
        expect(validBraces("{}[]()")).toBe(true)
    })
    test("()", () => {
        expect(validBraces("{}[]()")).toBe(true)
    })
    test("([{}])", () => {
        expect(validBraces("([{}])")).toBe(true)
    })
    test("(}", () => {
        expect(validBraces("(}")).toBe(false)
    })
    test("[({})](]", () => {
        expect(validBraces("[({})](]")).toBe(false)
    })

})
const {describe, test, expect} = require("@jest/globals");
const {removeDuplicates} = require("./removeDuplicates");


describe("remove duplicates", () => {
  test("first", () => {
    const arr = [0,0,1,1,1,2,2,3,3,4]
    expect(removeDuplicates(arr)).toBe(5)
  })
  test("empty", () => {
    const arr = []
    expect(removeDuplicates(arr)).toBe(0)
  })

})


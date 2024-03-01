const {describe, test, expect} = require("@jest/globals");
const { removeElementTwoPointers} = require("./removeElement");


describe("remove element two pointers", () =>{
  test("first", () => {
    const arr = [0,1,2,2,3,0,4,2,3]
    expect(removeElementTwoPointers(arr, 2)).toBe(6)
    console.log(arr)
  })
  test("zero value", () => {
    expect(removeElementTwoPointers([], 8)).toBe(0)
  })
})
import { describe, test, expect } from "@jest/globals";
import { removeDuplicates } from "./removeDuplicates";


describe("remove duplicates", () => {
  test("first", () => {
    const arr = [0,0,1,1,1,2,2,3,3,4]
    expect(removeDuplicates(arr)).toBe(5)
  })
  test("empty", () => {
    const arr: number[] = []
    expect(removeDuplicates(arr)).toBe(0)
  })
  test("not dublicate", () => {
    const arr = [1,2,3,4,5]
    expect(removeDuplicates(arr)).toBe(5)
  })

})


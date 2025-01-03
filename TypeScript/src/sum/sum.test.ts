import { sum } from './sum';
import { describe, test, expect } from "@jest/globals";

describe("first test", () => {
    test("1 + 1", () => {
        expect(sum(1,1)).toBe(2)
    })
})




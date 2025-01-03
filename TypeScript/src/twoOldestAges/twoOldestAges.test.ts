import { describe, test, expect } from "@jest/globals";
import { twoOldestAges } from './twoOldestAges';


describe("two oldest ages", () => {
    test("two element", () => {
        const ages = [98,45]
        expect(twoOldestAges(ages)).toEqual([98, 45])
    })
    test("empty array", () => {
        const ages:number[] = []
        expect(twoOldestAges(ages)).toEqual([])
    })
    test("success test", () => {
        const ages = [1, 5, 87, 45, 8, 8]
        expect(twoOldestAges(ages)).toEqual([45, 87])
    })
    // test("two old", () => {
    //     const ages = [1, 5, 87,87, 45, 8, 8]
    //     expect(twoOldestAges(ages)).toEqual([45, 87, 87])
    // })
})



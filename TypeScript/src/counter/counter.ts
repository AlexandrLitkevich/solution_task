

//https://leetcode.com/problems/counter/submissions/1459931798/
type nextFunc = () => number

export const Counter = function Counter(n: number): nextFunc {
    let localN = n;
    return function() {
        return localN++
    }
}
//TODO написать тест
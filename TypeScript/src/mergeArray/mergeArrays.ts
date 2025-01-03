

export const mergeArrays = (arr1: number[], arr2: number[]) : number[] =>  {
    const mergeMap = new Set(arr1.concat(arr2).sort((a,b) => a - b))
    return Array.from(mergeMap)
}

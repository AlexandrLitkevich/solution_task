

const mergeArrays = (arr1, arr2) =>  {
    const mergeMap = new Set(arr1.concat(arr2).sort((a,b) => a - b))
    return Array.from(mergeMap)
    //return [...new Set(a.concat(b))].sort((a,b)=>a-b)
}

module.exports = { mergeArrays}
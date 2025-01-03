

export const twoOldestAges = (ages: number[]): number[] => {
    if (ages.length < 2) return []

    if (ages.length === 2) return ages.sort((a, b) => b - a)

    return ages.sort((a, b) => b - a).slice(0, 2).reverse()
}
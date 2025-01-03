

export const validBraces = (braces: string): boolean => {
    const dictionary: Record<string, string> = {
        "(": ")",
        "{": "}",
        "[": "]",
    }

    var stack: string[] = []

    for(const char of braces) {
        if (dictionary[char]) {
            stack.push(char) 
        } else {
            const lastChar = stack.pop()
            if (lastChar === undefined || dictionary[lastChar] !== char) {
                return false; 
            } 
        }
    }

    return stack.length === 0
}
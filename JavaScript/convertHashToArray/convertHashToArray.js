

function convertHashToArray (hash) {
    if (hash === null) {
        return []
    }
    return  Object.entries(hash).sort()
}

module.exports = {convertHashToArray}
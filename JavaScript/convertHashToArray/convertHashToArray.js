

function convertHashToArray (hash) {
    if (hash === null) {
        return []
    }
    return Object.entries(hash)
}

module.exports = {convertHashToArray}
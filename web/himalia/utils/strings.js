const toCamelCase = (string) => {
    return string.replace(/(?:^\w|[A-Z]|\b\w)/g, function(word, index) {
        return index === 0 ? word.toUpperCase() : word.toLowerCase()
    }).replace(/\s+/g, '')
}

export { toCamelCase }

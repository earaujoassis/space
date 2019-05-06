const extractDataForm = (form, attributes) => {
    const root = form.getRootNode()
    let data = new FormData()

    for (let attr of attributes) {
        data.append(attr, root.getElementsByName(attr)[0].value)
    }

    return data
}

export { extractDataForm }

let SpaceApi = {
    createUser(data) {
        return fetch('/api/v1/users', {
            method: 'POST',
            body: data
        })
    }
}

export default SpaceApi

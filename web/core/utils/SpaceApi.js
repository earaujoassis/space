let SpaceApi = {
    createUser(data) {
        return fetch('/api/v1/users', {
            method: 'POST',
            body: data
        })
    },

    createSession(data) {
        return fetch('/api/v1/sessions', {
            method: 'POST',
            body: data
        })
    }
}

export default SpaceApi

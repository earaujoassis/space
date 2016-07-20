let SpaceApi = {
    createUser(data) {
        return fetch('/api/v1/users/create', {
            method: 'POST',
            body: data
        })
    },

    createSession(data) {
        return fetch('/api/v1/sessions/create', {
            method: 'POST',
            body: data
        })
    }
}

export default SpaceApi

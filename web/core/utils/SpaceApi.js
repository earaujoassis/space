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
    },

    fetchProfile(id, token) {
        return fetch(`/api/v1/users/${id}/profile`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`
            }
        })
    },

    fetchActiveClients(id, token) {
        return fetch(`/api/v1/users/${id}/clients`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`
            }
        })
    },

    revokeActiveClient(id, key, token) {
        return fetch(`/api/v1/users/${id}/clients/${key}/revoke`, {
            method: 'DELETE',
            headers: {
                Authorization: `Bearer ${token}`
            }
        })
    }
}

export default SpaceApi

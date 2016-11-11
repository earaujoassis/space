let SpaceApi = {
    createUser(data) {
        return fetch('/api/v1/users/create', {
            method: 'POST',
            headers: {
                'X-Requested-By': 'SpaceApi'
            },
            body: data
        })
    },

    fetchActiveClients(id, token) {
        return fetch(`/api/v1/users/${id}/clients`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi'
            }
        })
    },

    revokeActiveClient(id, key, token) {
        return fetch(`/api/v1/users/${id}/clients/${key}/revoke`, {
            method: 'DELETE',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi'
            }
        })
    },

    fetchProfile(id, token) {
        return fetch(`/api/v1/users/${id}/profile`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi'
            }
        })
    },

    createSession(data) {
        return fetch('/api/v1/sessions/create', {
            method: 'POST',
            headers: {
                'X-Requested-By': 'SpaceApi'
            },
            body: data
        })
    }
}

export default SpaceApi

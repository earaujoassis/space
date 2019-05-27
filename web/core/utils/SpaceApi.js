const SpaceApi = {
    createUser(data) {
        return fetch('/api/users/create', {
            method: 'POST',
            headers: {
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    },

    fetchActiveClients(id, token) {
        return fetch(`/api/users/${id}/clients`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            }
        })
    },

    revokeActiveClient(id, key, token) {
        return fetch(`/api/users/${id}/clients/${key}/revoke`, {
            method: 'DELETE',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            }
        })
    },

    fetchProfile(id, token) {
        return fetch(`/api/users/${id}/profile`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            }
        })
    },

    createSession(data) {
        return fetch('/api/sessions/create', {
            method: 'POST',
            headers: {
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    },

    adminify(id, token, data) {
        return fetch(`/api/users/${id}/adminify`, {
            method: 'PATCH',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    }
}

export default SpaceApi

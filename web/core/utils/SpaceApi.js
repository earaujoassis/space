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

    createClient(token, data) {
        return fetch('/api/clients/create', {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    },

    fetchClients(token) {
        return fetch('/api/clients', {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            }
        })
    }
}

export default SpaceApi

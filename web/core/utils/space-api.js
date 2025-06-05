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

    async updatePassword(data) {
        return await fetch('/api/users/update/password', {
            method: 'PATCH',
            headers: {
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    },

    async createSession(data) {
        return await fetch('/api/sessions/create', {
            method: 'POST',
            headers: {
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    },

    async createMagicSession(data) {
        return await fetch('/api/sessions/magic', {
            method: 'POST',
            headers: {
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    },

    async requestUpdate(data) {
        return await fetch('/api/users/update/request', {
            method: 'POST',
            headers: {
                'X-Requested-By': 'SpaceApi',
                Accept: 'application/vnd.space.v1+json'
            },
            body: data
        })
    }
}

export default SpaceApi

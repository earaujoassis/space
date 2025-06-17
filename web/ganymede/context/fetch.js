export const createSession = async (data) => {
    return await fetch('/api/sessions/create', {
        method: 'POST',
        headers: {
            'X-Requested-By': 'SpaceApi',
            Accept: 'application/vnd.space.v1+json'
        },
        body: data
    })
}

export const createMagicSession = async (data) => {
    return await fetch('/api/sessions/magic', {
        method: 'POST',
        headers: {
            'X-Requested-By': 'SpaceApi',
            Accept: 'application/vnd.space.v1+json'
        },
        body: data
    })
}

export const requestUpdate = async (data) => {
    return await fetch('/api/users/update/request', {
        method: 'POST',
        headers: {
            'X-Requested-By': 'SpaceApi',
            Accept: 'application/vnd.space.v1+json'
        },
        body: data
    })
}

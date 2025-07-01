export const updatePassword = async (data) => {
    return await fetch('/api/users/me/password', {
        method: 'PATCH',
        headers: {
            'X-Requested-By': 'SpaceApi',
            Accept: 'application/vnd.space.v1+json'
        },
        body: data
    })
}

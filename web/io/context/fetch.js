export const createUser = async data => {
  return await fetch('/api/users', {
    method: 'POST',
    headers: {
      'X-Requested-By': 'SpaceApi',
      Accept: 'application/vnd.space.v1+json',
    },
    body: data,
  })
}

const SpaceApi = {
  createUser (data) {
    return fetch('/api/users/create', {
      method: 'POST',
      headers: {
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  },

  requestUpdate (data) {
    return fetch('/api/users/update/request', {
      method: 'POST',
      headers: {
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  },

  updatePassword (data) {
    return fetch('/api/users/update/password', {
      method: 'PATCH',
      headers: {
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  },

  createSession (data) {
    return fetch('/api/sessions/create', {
      method: 'POST',
      headers: {
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  },

  createMagicSession (data) {
    return fetch('/api/sessions/magic', {
      method: 'POST',
      headers: {
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  }
};

export default SpaceApi;

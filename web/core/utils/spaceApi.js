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

  fetchProfile (id, token) {
    return fetch(`/api/users/${id}/profile`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      }
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

  adminify (token, data) {
    return fetch('/api/users/update/adminify', {
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${token}`,
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  },

  fetchActiveClients (id, token) {
    return fetch(`/api/users/${id}/clients`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      }
    });
  },

  revokeActiveClient (id, key, token) {
    return fetch(`/api/users/${id}/clients/${key}/revoke`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      }
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
  },

  fetchClients (token) {
    return fetch('/api/clients', {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      }
    });
  },

  createClient (token, data) {
    return fetch('/api/clients/create', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  },

  updateClient (id, token, data) {
    return fetch(`/api/clients/${id}/profile`, {
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${token}`,
        'X-Requested-By': 'SpaceApi',
        Accept: 'application/vnd.space.v1+json'
      },
      body: data
    });
  }
};

export default SpaceApi;

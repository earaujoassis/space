import axios from 'axios'

import store from '@stores'

axios.defaults.withCredentials = true

const fetch = axios.create({
  baseURL: '/api/',
  headers: {
    'X-Requested-By': 'SpaceApi',
    Accept: 'application/vnd.space.v1+json',
    'Content-Type': 'application/x-www-form-urlencoded',
  },
})

fetch.interceptors.request.use(config => {
  const state = store.getState()
  if (
    state &&
    state.root &&
    state.root.application &&
    state.root.application.action_token
  ) {
    const token = state.root.application.action_token
    config.headers['Authorization'] = `Bearer ${token}`
  }
  return config
})

export default fetch

/* eslint-disable quote-props */
import axios from 'axios'

axios.defaults.withCredentials = true

const fetch = axios.create({
    baseURL: '/api/',
    headers: {
        'X-Requested-By': 'SpaceApi',
        'Accept': 'application/vnd.space.v1+json',
        'Content-Type': 'application/x-www-form-urlencoded'
    }
})

export default fetch

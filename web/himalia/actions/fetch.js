/* eslint-disable quote-props */
import axios from 'axios'

const fetch = axios.create({
    baseURL: '/api/',
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
    }
})

export default fetch

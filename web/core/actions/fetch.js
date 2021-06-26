/* eslint-disable quote-props */
import axios from 'axios';

const fetch = axios.create({
  baseURL: '/api/',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/vnd.space.v1+json',
    'X-Requested-By': 'SpaceApi'
  }
});

export default fetch;

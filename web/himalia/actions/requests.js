import {
  USER_REQUEST_START,
  USER_REQUEST_SUCCESS,
  USER_REQUEST_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const userRequestStart = () => {
  return {
    type: USER_REQUEST_START,
  }
}

export const userRequestSuccess = () => {
  return {
    type: USER_REQUEST_SUCCESS,
  }
}

export const userRequestError = error => {
  return {
    type: USER_REQUEST_ERROR,
    error: error,
  }
}

export const requestEmailVerification = (holder, email) => {
  return dispatch => {
    const data = new FormData()
    data.append('request_type', 'email_verification')
    data.append('holder', holder)
    data.append('email', email)
    dispatch(userRequestStart())
    fetch
      .post('users/me/requests', data)
      .then(response => {
        dispatch(userRequestSuccess(response.data))
      })
      .catch(error => {
        dispatch(userRequestError(error))
        dispatch(toastError(error))
      })
  }
}

export const requestResetPassword = username => {
  return dispatch => {
    const data = new FormData()
    data.append('request_type', 'password')
    data.append('holder', username)
    dispatch(userRequestStart())
    fetch
      .post('users/me/requests', data)
      .then(response => {
        dispatch(userRequestSuccess(response.data))
      })
      .catch(error => {
        dispatch(userRequestError(error))
        dispatch(toastError(error))
      })
  }
}

export const requestResetSecretCodes = username => {
  return dispatch => {
    const data = new FormData()
    data.append('request_type', 'secrets')
    data.append('holder', username)
    dispatch(userRequestStart())
    fetch
      .post('users/me/requests', data)
      .then(response => {
        dispatch(userRequestSuccess(response.data))
      })
      .catch(error => {
        dispatch(userRequestError(error))
        dispatch(toastError(error))
      })
  }
}

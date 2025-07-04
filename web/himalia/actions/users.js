import * as actionTypes from './types'
import fetch from './fetch'

export const userRecordStart = () => {
  return {
    type: actionTypes.USER_RECORD_START,
  }
}

export const userRecordSuccess = data => {
  return {
    type: actionTypes.USER_RECORD_SUCCESS,
    user: data.user,
  }
}

export const userRecordError = error => {
  return {
    type: actionTypes.USER_RECORD_ERROR,
    error: error,
  }
}

export const userRequestStart = () => {
  return {
    type: actionTypes.USER_REQUEST_START,
  }
}

export const userRequestSuccess = () => {
  return {
    type: actionTypes.USER_REQUEST_SUCCESS,
  }
}

export const userRequestError = error => {
  return {
    type: actionTypes.USER_REQUEST_ERROR,
    error: error,
  }
}

export const fetchUserProfile = id => {
  return dispatch => {
    dispatch(userRecordStart())
    fetch
      .get(`users/${id}/profile`)
      .then(response => {
        dispatch(userRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(userRecordError(error))
      })
  }
}

export const becomeAdmin = (id, key) => {
  return dispatch => {
    const data = new FormData()
    data.append('user_id', id)
    data.append('application_key', key)
    dispatch(userRecordStart())
    fetch
      .patch('users/me/admin', data)
      .then(response => {
        dispatch(userRecordSuccess(response.data))
        window.location.reload()
      })
      .catch(error => {
        dispatch(userRecordError(error))
      })
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
      })
  }
}

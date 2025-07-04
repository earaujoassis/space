import {
  USER_RECORD_START,
  USER_RECORD_SUCCESS,
  USER_RECORD_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const userRecordStart = () => {
  return {
    type: USER_RECORD_START,
  }
}

export const userRecordSuccess = data => {
  return {
    type: USER_RECORD_SUCCESS,
    user: data.user,
  }
}

export const userRecordError = error => {
  return {
    type: USER_RECORD_ERROR,
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
        dispatch(toastError(error))
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
        dispatch(toastError(error))
      })
  }
}

import {
  EMAIL_RECORD_START,
  EMAIL_RECORD_SUCCESS,
  EMAIL_RECORD_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const emailRecordStart = () => {
  return {
    type: EMAIL_RECORD_START,
  }
}

export const emailRecordSuccess = (data, status) => {
  return {
    type: EMAIL_RECORD_SUCCESS,
    emails: data ? data.emails : undefined,
    stale: status == 204,
  }
}

export const emailRecordError = error => {
  return {
    type: EMAIL_RECORD_ERROR,
    error: error,
  }
}

export const fetchEmails = () => {
  return dispatch => {
    dispatch(emailRecordStart())
    fetch
      .get('users/me/emails')
      .then(response => {
        dispatch(emailRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(emailRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const addEmail = data => {
  return dispatch => {
    dispatch(emailRecordStart())
    fetch
      .post('users/me/emails', data)
      .then(response => {
        dispatch(emailRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(emailRecordError(error))
        dispatch(toastError(error))
      })
  }
}

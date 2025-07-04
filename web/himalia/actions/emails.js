import * as actionTypes from './types'
import fetch from './fetch'

export const emailRecordStart = () => {
  return {
    type: actionTypes.EMAIL_RECORD_START,
  }
}

export const emailRecordSuccess = data => {
  return {
    type: actionTypes.EMAIL_RECORD_SUCCESS,
    emails: data.emails,
  }
}

export const emailRecordError = error => {
  return {
    type: actionTypes.EMAIL_RECORD_ERROR,
    error: error,
  }
}

export const fetchEmails = () => {
  return dispatch => {
    dispatch(emailRecordStart())
    fetch
      .get('users/me/emails')
      .then(response => {
        dispatch(emailRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(emailRecordError(error))
      })
  }
}

export const addEmail = data => {
  return dispatch => {
    dispatch(emailRecordStart())
    fetch
      .post('users/me/emails', data)
      .then(() => {
        dispatch(emailRecordSuccess())
      })
      .catch(error => {
        dispatch(emailRecordError(error))
      })
  }
}

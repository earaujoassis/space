import {
  SESSION_RECORD_START,
  SESSION_RECORD_SUCCESS,
  SESSION_RECORD_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const sessionRecordStart = () => {
  return {
    type: SESSION_RECORD_START,
  }
}

export const sessionRecordSuccess = (data, status) => {
  return {
    type: SESSION_RECORD_SUCCESS,
    sessions: data ? data.sessions : undefined,
    stale: status == 204,
  }
}

export const sessionRecordError = error => {
  return {
    type: SESSION_RECORD_ERROR,
    error: error,
  }
}

export const fetchApplicationSessionsForUser = id => {
  return dispatch => {
    dispatch(sessionRecordStart())
    fetch
      .get(`users/${id}/sessions`)
      .then(response => {
        dispatch(sessionRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(sessionRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const revokeApplicationSessionForUser = (userId, sessionId) => {
  return dispatch => {
    dispatch(sessionRecordStart())
    fetch
      .delete(`users/${userId}/sessions/${sessionId}/revoke`)
      .then(response => {
        dispatch(sessionRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(sessionRecordError(error))
        dispatch(toastError(error))
      })
  }
}

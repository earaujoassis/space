import {
  APPLICATION_RECORD_START,
  APPLICATION_RECORD_SUCCESS,
  APPLICATION_RECORD_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const applicationRecordStart = () => {
  return {
    type: APPLICATION_RECORD_START,
  }
}

export const applicationRecordSuccess = (data, status) => {
  return {
    type: APPLICATION_RECORD_SUCCESS,
    clients: data ? data.clients : undefined,
    stale: status == 204,
  }
}

export const applicationRecordError = error => {
  return {
    type: APPLICATION_RECORD_ERROR,
    error: error,
  }
}

export const fetchClientApplicationsFromUser = id => {
  return dispatch => {
    dispatch(applicationRecordStart())
    fetch
      .get(`users/${id}/clients`)
      .then(response => {
        dispatch(applicationRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(applicationRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const revokeClientApplicationFromUser = (userId, clientId) => {
  return dispatch => {
    dispatch(applicationRecordStart())
    fetch
      .delete(`users/${userId}/clients/${clientId}/revoke`)
      .then(response => {
        dispatch(applicationRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(applicationRecordError(error))
        dispatch(toastError(error))
      })
  }
}

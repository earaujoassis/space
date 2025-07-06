import {
  SERVICE_RECORD_START,
  SERVICE_RECORD_SUCCESS,
  SERVICE_RECORD_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const serviceRecordStart = () => {
  return {
    type: SERVICE_RECORD_START,
  }
}

export const serviceRecordSuccess = (data, status) => {
  return {
    type: SERVICE_RECORD_SUCCESS,
    services: data ? data.services : undefined,
    stale: status == 204,
  }
}

export const serviceRecordError = error => {
  return {
    type: SERVICE_RECORD_ERROR,
    error: error,
  }
}

export const createService = data => {
  return dispatch => {
    dispatch(serviceRecordStart())
    fetch
      .post('services', data)
      .then(response => {
        dispatch(serviceRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(serviceRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const fetchServices = () => {
  return dispatch => {
    dispatch(serviceRecordStart())
    fetch
      .get('services')
      .then(response => {
        dispatch(serviceRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(serviceRecordError(error))
        dispatch(toastError(error))
      })
  }
}

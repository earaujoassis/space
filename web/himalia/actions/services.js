import * as actionTypes from './types'
import fetch from './fetch'

export const serviceRecordStart = () => {
  return {
    type: actionTypes.SERVICE_RECORD_START,
  }
}

export const serviceRecordSuccess = data => {
  return {
    type: actionTypes.SERVICE_RECORD_SUCCESS,
    services: data.services,
  }
}

export const serviceRecordError = error => {
  return {
    type: actionTypes.SERVICE_RECORD_ERROR,
    error: error,
  }
}

export const createService = data => {
  return dispatch => {
    dispatch(serviceRecordStart())
    fetch
      .post('services', data)
      .then(response => {
        dispatch(serviceRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(serviceRecordError(error))
      })
  }
}

export const fetchServices = () => {
  return dispatch => {
    dispatch(serviceRecordStart())
    fetch
      .get('services')
      .then(response => {
        dispatch(serviceRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(serviceRecordError(error))
      })
  }
}

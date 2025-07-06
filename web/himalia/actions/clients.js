import {
  CLIENT_RECORD_START,
  CLIENT_RECORD_SUCCESS,
  CLIENT_RECORD_ERROR,
  CLIENT_RECORD_STALE,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const clientRecordStart = () => {
  return {
    type: CLIENT_RECORD_START,
  }
}

export const clientRecordSuccess = (data, status) => {
  return {
    type: CLIENT_RECORD_SUCCESS,
    clients: data ? data.clients : undefined,
    stale: status == 204,
  }
}

export const clientRecordError = error => {
  return {
    type: CLIENT_RECORD_ERROR,
    error: error,
  }
}

export const clientRecordStale = () => {
  return {
    type: CLIENT_RECORD_STALE
  }
}

export const createClient = data => {
  return dispatch => {
    dispatch(clientRecordStart())
    fetch
      .post('clients', data)
      .then(response => {
        dispatch(clientRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const fetchClients = () => {
  return dispatch => {
    dispatch(clientRecordStart())
    fetch
      .get('clients')
      .then(response => {
        dispatch(clientRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const setClientForEdition = client => {
  return dispatch => {
    dispatch(clientRecordSuccess({ clients: [client] }))
  }
}

export const updateClient = (id, data) => {
  return dispatch => {
    dispatch(clientRecordStart())
    fetch
      .patch(`clients/${id}/profile`, data)
      .then(response => {
        dispatch(clientRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const staleClientRecords = () => {
  return dispatch => {
    dispatch(clientRecordStale())
  }
}

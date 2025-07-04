import * as actionTypes from './types'
import fetch from './fetch'

export const clientRecordStart = () => {
  return {
    type: actionTypes.CLIENT_RECORD_START,
  }
}

export const clientRecordSuccess = data => {
  return {
    type: actionTypes.CLIENT_RECORD_SUCCESS,
    clients: data.clients,
  }
}

export const clientRecordError = error => {
  return {
    type: actionTypes.CLIENT_RECORD_ERROR,
    error: error,
  }
}

export const createClient = data => {
  return dispatch => {
    dispatch(clientRecordStart())
    fetch
      .post('clients', data)
      .then(response => {
        dispatch(clientRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
      })
  }
}

export const fetchClients = () => {
  return dispatch => {
    dispatch(clientRecordStart())
    fetch
      .get('clients')
      .then(response => {
        dispatch(clientRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
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
        dispatch(clientRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
      })
  }
}

export const fetchClientApplicationsFromUser = id => {
  return dispatch => {
    dispatch(clientRecordStart())
    fetch
      .get(`users/${id}/clients`)
      .then(response => {
        dispatch(clientRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
      })
  }
}

export const revokeClientApplicationFromUser = (userId, clientId) => {
  return dispatch => {
    dispatch(clientRecordStart())
    fetch
      .delete(`users/${userId}/clients/${clientId}/revoke`)
      .then(response => {
        dispatch(clientRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(clientRecordError(error))
      })
  }
}

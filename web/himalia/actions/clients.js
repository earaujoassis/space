import * as actionTypes from './types'
import fetch from './fetch'

export const clientRecordStart = () => {
    return {
        type: actionTypes.CLIENT_RECORD_START
    }
}

export const clientRecordSuccess = (data) => {
    return {
        type: actionTypes.CLIENT_RECORD_SUCCESS,
        clients: data.clients
    }
}

export const clientRecordError = (error) => {
    return {
        type: actionTypes.CLIENT_RECORD_ERROR,
        error: error
    }
}

export const createClient = (data, token) => {
    return dispatch => {
        dispatch(clientRecordStart())
        fetch.post('clients/create', data, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(clientRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(clientRecordError(error))
            })
    }
}

export const fetchClients = (token) => {
    return dispatch => {
        dispatch(clientRecordStart())
        fetch.get('clients', { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(clientRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(clientRecordError(error))
            })
    }
}

export const setClientForEdition = (client) => {
    return dispatch => {
        dispatch(clientRecordSuccess({ clients: [client] }))
    }
}

export const updateClient = (id, data, token) => {
    return dispatch => {
        dispatch(clientRecordStart())
        fetch.patch(`clients/${id}/profile`, data, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(clientRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(clientRecordError(error))
            })
    }
}

export const fetchClientApplicationsFromUser = (id, token) => {
    return dispatch => {
        dispatch(clientRecordStart())
        fetch.get(`users/${id}/clients`, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(clientRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(clientRecordError(error))
            })
    }
}

export const revokeClientApplicationFromUser = (userId, clientId, token) => {
    return dispatch => {
        dispatch(clientRecordStart())
        fetch.delete(`users/${userId}/clients/${clientId}/revoke`, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(clientRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(clientRecordError(error))
            })
    }
}

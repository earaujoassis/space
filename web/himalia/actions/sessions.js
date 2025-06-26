import * as actionTypes from './types'
import fetch from './fetch'

export const sessionRecordStart = () => {
    return {
        type: actionTypes.SESSION_RECORD_START
    }
}

export const sessionRecordSuccess = (data) => {
    return {
        type: actionTypes.SESSION_RECORD_SUCCESS,
        sessions: data.sessions
    }
}

export const sessionRecordError = (error) => {
    return {
        type: actionTypes.SESSION_RECORD_ERROR,
        error: error
    }
}

export const fetchApplicationSessionsForUser = (id, token) => {
    return dispatch => {
        dispatch(sessionRecordStart())
        fetch.get(`users/${id}/sessions`, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(sessionRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(sessionRecordError(error))
            })
    }
}

export const revokeApplicationSessionForUser = (userId, sessionId, token) => {
    return dispatch => {
        dispatch(sessionRecordStart())
        fetch.delete(`users/${userId}/sessions/${sessionId}/revoke`, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(sessionRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(sessionRecordError(error))
            })
    }
}

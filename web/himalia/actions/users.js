import * as actionTypes from './types'
import fetch from './fetch'

export const userRecordStart = () => {
    return {
        type: actionTypes.USER_RECORD_START
    }
}

export const userRecordSuccess = (data) => {
    return {
        type: actionTypes.USER_RECORD_SUCCESS,
        user: data.user
    }
}

export const userRecordError = (error) => {
    return {
        type: actionTypes.USER_RECORD_ERROR,
        error: error
    }
}

export const fetchUserProfile = (id, token) => {
    return dispatch => {
        dispatch(userRecordStart())
        fetch.get(`users/${id}/profile`, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(userRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(userRecordError(error))
            })
    }
}

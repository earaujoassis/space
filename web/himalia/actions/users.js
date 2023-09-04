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

export const adminifyUser = (id, key, token) => {
    return dispatch => {
        const data = new FormData()
        data.append('user_id', id)
        data.append('application_key', key)
        dispatch(userRecordStart())
        fetch.patch('users/update/adminify', data, { headers: { 'Authorization': `Bearer ${token}` } })
            .then(response => {
                dispatch(userRecordSuccess(response.data))
                window.location.reload()
            })
            .catch(error => {
                dispatch(userRecordError(error))
            })
    }
}

export const requestResetPassword = (username) => {
    return dispatch => {
        const data = new FormData()
        data.append('request_type', 'password')
        data.append('holder', username)
        dispatch(userRecordStart())
        fetch.post('users/update/request', data)
            .then(response => {
                dispatch(userRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(userRecordError(error))
            })
    }
}

export const requestResetSecretCodes = (username) => {
    return dispatch => {
        const data = new FormData()
        data.append('request_type', 'secrets')
        data.append('holder', username)
        dispatch(userRecordStart())
        fetch.post('users/update/request', data)
            .then(response => {
                dispatch(userRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(userRecordError(error))
            })
    }
}

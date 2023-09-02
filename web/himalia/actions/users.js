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

export const fetchUser = () => {
    return dispatch => {
        dispatch(userRecordStart())
        fetch.get('users/')
            .then(response => {
                dispatch(userRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(userRecordError(error))
            })
    }
}

export const updateUser = (id, data) => {
    return dispatch => {
        dispatch(userRecordStart())
        fetch.patch(`users/${id}`, data)
            .then(response => {
                dispatch(userRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(userRecordError(error))
            })
    }
}

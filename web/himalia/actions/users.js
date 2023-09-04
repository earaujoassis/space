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

export const fetchUser = (id, token) => {
    return dispatch => {
        dispatch(userRecordStart())
        fetch.get(`users/${id}/profile`, {}, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'X-Requested-By': 'SpaceApi',
                'Accept': 'application/vnd.space.v1+json',
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        })
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

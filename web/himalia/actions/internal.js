import * as actionTypes from './types'
import fetch from './fetch'

export const internalSetToastDisplay = (displayToast = false) => {
    return {
        type: actionTypes.INTERNAL_DISPLAY_TOAST,
        displayToast
    }
}

export const internalRecordStart = () => {
    return {
        type: actionTypes.INTERNAL_RECORD_START
    }
}

export const internalRecordSuccess = (data) => {
    return {
        type: actionTypes.INTERNAL_RECORD_SUCCESS,
        application: data.application
    }
}

export const internalRecordError = (error) => {
    return {
        type: actionTypes.INTERNAL_RECORD_ERROR,
        error: error
    }
}

export const fetchWorkspace = () => {
    return dispatch => {
        dispatch(internalRecordStart())
        fetch.get('users/me/workspace')
            .then(response => {
                dispatch(internalRecordSuccess(response.data))
            })
            .catch(error => {
                dispatch(internalRecordError(error))
            })
    }
}

import {
  SETTING_RECORD_START,
  SETTING_RECORD_SUCCESS,
  SETTING_RECORD_ERROR,
} from './types'
import fetch from './fetch'
import { toastError } from './internal'

export const settingRecordStart = () => {
  return {
    type: SETTING_RECORD_START,
  }
}

export const settingRecordSuccess = (data, status) => {
  return {
    type: SETTING_RECORD_SUCCESS,
    settings: data ? data.settings : undefined,
    stale: status == 204,
  }
}

export const settingRecordError = error => {
  return {
    type: SETTING_RECORD_ERROR,
    error: error,
  }
}

export const fetchUserSettings = () => {
  return dispatch => {
    dispatch(settingRecordStart())
    fetch
      .get(`users/me/settings`)
      .then(response => {
        dispatch(settingRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(settingRecordError(error))
        dispatch(toastError(error))
      })
  }
}

export const patchUserSettings = data => {
  return dispatch => {
    dispatch(settingRecordStart())
    fetch
      .patch(`users/me/settings`, data)
      .then(response => {
        dispatch(settingRecordSuccess(response.data, response.status))
      })
      .catch(error => {
        dispatch(settingRecordError(error))
        dispatch(toastError(error))
      })
  }
}

import * as actionTypes from './types'
import fetch from './fetch'

export const settingRecordStart = () => {
  return {
    type: actionTypes.SETTING_RECORD_START,
  }
}

export const settingRecordSuccess = data => {
  return {
    type: actionTypes.SETTING_RECORD_SUCCESS,
    settings: data.settings,
  }
}

export const settingRecordError = error => {
  return {
    type: actionTypes.SETTING_RECORD_ERROR,
    error: error,
  }
}

export const fetchUserSettings = () => {
  return dispatch => {
    dispatch(settingRecordStart())
    fetch
      .get(`users/me/settings`)
      .then(response => {
        dispatch(settingRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(settingRecordError(error))
      })
  }
}

export const patchUserSettings = data => {
  return dispatch => {
    dispatch(settingRecordStart())
    fetch
      .patch(`users/me/settings`, data)
      .then(response => {
        dispatch(settingRecordSuccess(response.data))
      })
      .catch(error => {
        dispatch(settingRecordError(error))
      })
  }
}

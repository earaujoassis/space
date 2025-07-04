import {
  SETTING_RECORD_START,
  SETTING_RECORD_SUCCESS,
  SETTING_RECORD_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const settingRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const settingRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.settings || state.data,
    stale: action.stale,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const settingRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case SETTING_RECORD_START:
      return settingRecordStart(state, action)
    case SETTING_RECORD_SUCCESS:
      return settingRecordSuccess(state, action)
    case SETTING_RECORD_ERROR:
      return settingRecordError(state, action)
    default:
      return state
  }
}

export default reducer

import {
  EMAIL_RECORD_START,
  EMAIL_RECORD_SUCCESS,
  EMAIL_RECORD_ERROR,
} from '@actions/types'

import resourceTemplate from './template'

const initialState = { ...resourceTemplate }

const emailRecordStart = state => {
  NProgress.start()
  const updateState = { loading: true }
  return { ...state, ...updateState }
}

const emailRecordSuccess = (state, action) => {
  NProgress.done()
  const updateState = {
    data: action.emails || state.data,
    stale: action.stale,
    loading: false,
    error: null,
  }
  return { ...state, ...updateState }
}

const emailRecordError = (state, action) => {
  NProgress.done()
  const updateState = {
    loading: false,
    error: action.error,
  }
  return { ...state, ...updateState }
}

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case EMAIL_RECORD_START:
      return emailRecordStart(state, action)
    case EMAIL_RECORD_SUCCESS:
      return emailRecordSuccess(state, action)
    case EMAIL_RECORD_ERROR:
      return emailRecordError(state, action)
    default:
      return state
  }
}

export default reducer
